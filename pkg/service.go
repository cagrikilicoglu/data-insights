package pkg

import (
	"data-insights/kit/ai"
	"data-insights/kit/email"
	"data-insights/kit/file"
	"data-insights/kit/metrics"
	"data-insights/kit/model"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
)

func ProcessFiles(envVariables model.EnvVariables) error {
	files, err := file.FilePathWalkDir(envVariables.FileDirectory)
	if err != nil {
		return fmt.Errorf("error loading files: %s from directory: %s\n", err, envVariables.FileDirectory)
	}

	for _, fileSource := range files {
		err = processFile(envVariables, fileSource)
		if err != nil {
			return fmt.Errorf("error processing file %s: %s\n", fileSource, err)
		}
		log.Printf("Email sent successfully, insights from the file %s has been sent to %s\n", fileSource, envVariables.EmailTo)
	}
	return nil
}

func processFile(envVariables model.EnvVariables, fileSource string) error {

	data, err := file.GetRawDataFromFile(fileSource)
	if err != nil {
		return fmt.Errorf("error while getting raw data from file %s: %v\n", fileSource, err)
	}

	openAiClient := ai.NewOpenAIClient(ai.GPT4oMini, ai.OpenAIUrl, ai.OpenAIMaxTokens, ai.OpenAISenderRole)
	insights, err := openAiClient.GetInsightsFromLLM(envVariables.ApiKey, metrics.GetImportantMetrics(data))
	if err != nil {
		return fmt.Errorf("error getting insights from LLM: %v\n", err)
	}

	var userMetricsWithInsights model.UserMetricsWithInsights
	err = json.Unmarshal([]byte(insights), &userMetricsWithInsights)
	if err != nil {
		return fmt.Errorf("error unmarshalling user metrics with insights from LLM: %v\n", err)
	}

	emailService := email.NewSMTPEmailService(envVariables.SmtpHost, envVariables.SmtpPort, envVariables.EmailFrom, envVariables.EmailPass)

	// Set up the renderer
	renderer := email.NewRenderer(filepath.Join(email.TemplateDir, email.TemplateFile))
	// Render the email body
	body, err := renderer.Render(model.EmailData{
		RecipientName:           envVariables.RecipientName,
		UserMetricsWithInsights: userMetricsWithInsights,
	})
	if err != nil {
		return fmt.Errorf("error rendering email template: %v", err)
	}

	// Send the email
	err = emailService.SendEmail(envVariables.EmailTo, email.SubjectName, body)
	if err != nil {
		return fmt.Errorf("error sending email to %s: %v", envVariables.EmailTo, err)
	}

	return nil
}

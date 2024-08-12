package pkg

import (
	"data-insights/kit/ai"
	"data-insights/kit/common"
	"data-insights/kit/email"
	"data-insights/kit/file"
	"data-insights/kit/metrics"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
)

// ProcessFiles iterates over all files in the specified directory, processes each file to generate insights,
// and sends an email with the insights. Returns an error if any step fails.
func ProcessFiles(envVariables common.EnvVariables, openAiClient ai.OpenAIClient) error {

	files, err := file.FilePathWalkDir(envVariables.FileDirectory)
	if err != nil {
		return fmt.Errorf("error loading files: %s from directory: %s", err, envVariables.FileDirectory)
	}

	for _, fileSource := range files {
		err = processFile(envVariables, fileSource, openAiClient)
		if err != nil {
			return fmt.Errorf("error processing file %s: %s", fileSource, err)
		}
		log.Printf("Email sent successfully, insights from the file %s has been sent to %s", fileSource, envVariables.EmailTo)
	}
	return nil
}

// processFile handles the processing of a single file. It reads the raw data from the file,
// generates insights using OpenAI, unmarshalls the results, renders an email template with the insights,
// and sends the email. Returns an error if any step fails.
func processFile(envVariables common.EnvVariables, fileSource string, openAiClient ai.OpenAIClient) error {

	data, err := file.GetRawDataFromFile(fileSource)
	if err != nil {
		return fmt.Errorf("error while getting raw data from file %s: %v", fileSource, err)
	}

	insights, err := openAiClient.GetInsightsFromLLM(envVariables.ApiKey, metrics.CalculateKeyMetrics(data))
	if err != nil {
		return fmt.Errorf("error getting insights from LLM: %v", err)
	}

	var userMetricsWithInsights common.UserMetricsWithInsights
	err = json.Unmarshal([]byte(insights), &userMetricsWithInsights)
	if err != nil {
		return fmt.Errorf("error unmarshalling user metrics with insights from LLM: %v", err)
	}

	emailService := email.NewSMTPEmailService(envVariables.SmtpHost, envVariables.SmtpPort, envVariables.EmailFrom, envVariables.EmailPass)

	// Set up the renderer
	renderer := email.NewRenderer(filepath.Join(email.TemplateDir, email.TemplateFile))
	// Render the email body
	body, err := renderer.Render(common.EmailData{
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

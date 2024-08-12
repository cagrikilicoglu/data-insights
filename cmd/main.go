package main

import (
	"data-insights/kit/ai"
	"data-insights/kit/email"
	"data-insights/kit/file"
	"data-insights/kit/metrics"
	"data-insights/kit/model"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

const thresholdDataPointNumber int = 100

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	fileDirectory := os.Getenv("FILE_DIR")
	if fileDirectory == "" {
		log.Fatal("There is no file path supplied")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is required")
	}

	emailFrom := os.Getenv("EMAIL_FROM")
	if emailFrom == "" {
		log.Fatal("EMAIL_FROM environment variable is required")
	}

	emailFromPass := os.Getenv("EMAIL_FROM_PASS")
	if emailFromPass == "" {
		log.Fatal("EMAIL_FROM_PASS environment variable is required")
	}

	emailTo := os.Getenv("EMAIL_TO")
	if emailTo == "" {
		log.Fatal("EMAIL_TO environment variable is required")
	}

	smtpHost := os.Getenv("SMTP_HOST")
	if smtpHost == "" {
		log.Fatal("SMTP_HOST environment variable is required")
	}

	smtpPort := os.Getenv("SMTP_PORT")
	if smtpPort == "" {
		log.Fatal("SMTP_PORT environment variable is required")
	}

	files, err := file.FilePathWalkDir(fileDirectory)
	if err != nil {
		panic(err)
	}

	for _, fileSource := range files {
		data, err := file.ProcessJSONFile(fileSource)
		if err != nil {
			log.Printf("Error processing file %s: %v\n", fileSource, err)
			continue
		}
		importantMetrics := metrics.GetImportantMetrics(data)
		prompt := ai.CreatePrompt(importantMetrics)
		fmt.Println(prompt)

		insights, err := ai.GetInsightsFromLLM(apiKey, importantMetrics)
		if err != nil {
			log.Printf("Error getting insights from LLM: %v\n", err)
			continue
		}

		var userMetricsWithInsights model.UserMetricsWithInsights
		err = json.Unmarshal([]byte(insights), &userMetricsWithInsights)
		if err != nil {
			log.Printf("error when unmarshalling")
			return
		}

		fmt.Printf("Insights for file %s: %s\n", fileSource, insights)
		fmt.Println("User Insights for file %s: %s\n", fileSource, userMetricsWithInsights)

		emailService := email.NewSMTPEmailService(smtpHost, smtpPort, emailFrom, emailFromPass)

		// Set up the renderer
		renderer := email.NewRenderer(filepath.Join("templates", "email_template.html"))

		// Render the email body
		body, err := renderer.Render(model.EmailData{
			RecipientName:           "John Doe",
			UserMetricsWithInsights: userMetricsWithInsights,
		})
		if err != nil {
			log.Fatalf("Failed to render email template: %v", err)
		}

		// Send the email
		err = emailService.SendEmail(emailTo, "Website Insights Report", body)
		if err != nil {
			log.Fatalf("Failed to send email: %v", err)
		}

		log.Println("Email sent successfully")
	}
}

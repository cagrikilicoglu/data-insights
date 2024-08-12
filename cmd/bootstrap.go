package main

import (
	"data-insights/kit/model"
	"data-insights/pkg"
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Bootstrap() {
	envVariables, err := getEnvVariables()
	if err != nil {
		log.Fatalf("error loading environment variables: %s", err)
	}

	err = pkg.ProcessFiles(envVariables)
	if err != nil {
		log.Fatalf("error processing files: %s", err)
	}
}

func getEnvVariables() (model.EnvVariables, error) {

	if err := godotenv.Load(); err != nil {
		return model.EnvVariables{}, errors.New("no .env file found")
	}

	fileDirectory := os.Getenv("FILE_DIR")
	if fileDirectory == "" {
		return model.EnvVariables{}, errors.New("FILE_DIR environment variable is required")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return model.EnvVariables{}, errors.New("OPENAI_API_KEY environment variable is required")
	}

	emailFrom := os.Getenv("EMAIL_FROM")
	if emailFrom == "" {
		return model.EnvVariables{}, errors.New("EMAIL_FROM environment variable is required")
	}

	emailFromPass := os.Getenv("EMAIL_FROM_PASS")
	if emailFromPass == "" {
		return model.EnvVariables{}, errors.New("EMAIL_FROM_PASS environment variable is required")
	}

	emailTo := os.Getenv("EMAIL_TO")
	if emailTo == "" {
		return model.EnvVariables{}, errors.New("EMAIL_TO environment variable is required")
	}

	recipientName := os.Getenv("RECIPIENT_NAME")
	if emailTo == "" {
		return model.EnvVariables{}, errors.New("RECIPIENT_NAME environment variable is required")
	}

	smtpHost := os.Getenv("SMTP_HOST")
	if smtpHost == "" {
		return model.EnvVariables{}, errors.New("SMTP_HOST environment variable is required")
	}

	smtpPort := os.Getenv("SMTP_PORT")
	if smtpPort == "" {
		return model.EnvVariables{}, errors.New("SMTP_PORT environment variable is required")
	}

	return model.EnvVariables{
		FileDirectory: fileDirectory,
		ApiKey:        apiKey,
		EmailFrom:     emailFrom,
		EmailPass:     emailFromPass,
		EmailTo:       emailTo,
		RecipientName: recipientName,
		SmtpHost:      smtpHost,
		SmtpPort:      smtpPort,
	}, nil
}

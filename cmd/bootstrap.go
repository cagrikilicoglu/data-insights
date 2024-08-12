package main

import (
	"data-insights/kit/ai"
	"data-insights/kit/common"
	"data-insights/pkg"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

// Bootstrap initializes the application by loading environment variables and processing the files.
// It logs a fatal error and terminates the application if any of the critical steps fail.
func Bootstrap() {
	envVariables, err := getEnvVariables()
	if err != nil {
		log.Fatalf("error loading environment variables: %s", err)
	}

	openAiClient := ai.NewOpenAIClient(ai.GPT4oMini, ai.OpenAIUrl, ai.OpenAIMaxTokens, ai.OpenAISenderRole, &http.Client{})
	err = pkg.ProcessFiles(envVariables, openAiClient)
	if err != nil {
		log.Fatalf("error processing files: %s", err)
	}
}

// getEnvVariables loads and validates the required environment variables from a `.env` file.
// It uses a map to streamline the process of checking for each variable.
// Returns a populated EnvVariables struct if all required variables are set, or an error if any are missing.
func getEnvVariables() (common.EnvVariables, error) {

	if err := godotenv.Load(); err != nil {
		return common.EnvVariables{}, errors.New("no .env file found")
	}

	vars := map[string]string{
		"FILE_DIR":        "",
		"OPENAI_API_KEY":  "",
		"EMAIL_FROM":      "",
		"EMAIL_FROM_PASS": "",
		"EMAIL_TO":        "",
		"RECIPIENT_NAME":  "",
		"SMTP_HOST":       "",
		"SMTP_PORT":       "",
	}

	for key := range vars {
		vars[key] = os.Getenv(key)
		if vars[key] == "" {
			return common.EnvVariables{}, fmt.Errorf("%s environment variable is required", key)
		}
	}

	return common.EnvVariables{
		FileDirectory: vars["FILE_DIR"],
		ApiKey:        vars["OPENAI_API_KEY"],
		EmailFrom:     vars["EMAIL_FROM"],
		EmailPass:     vars["EMAIL_FROM_PASS"],
		EmailTo:       vars["EMAIL_TO"],
		RecipientName: vars["RECIPIENT_NAME"],
		SmtpHost:      vars["SMTP_HOST"],
		SmtpPort:      vars["SMTP_PORT"],
	}, nil
}

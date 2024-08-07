package main

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

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

	files, err := FilePathWalkDir(fileDirectory)
	if err != nil {
		panic(err)
	}

	for _, fileSource := range files {
		data, err := processJSONFile(fileSource)
		if err != nil {
			log.Printf("Error processing file %s: %v\n", fileSource, err)
			continue
		}
		insights, err := getInsightsFromLLM(apiKey, data)
		if err != nil {
			log.Printf("Error getting insights from LLM: %v\n", err)
			continue
		}
		fmt.Printf("Insights for file %s: %s\n", fileSource, insights)
	}
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// processJSONFile reads and parses the JSON file
func processJSONFile(filePath string) ([]Insight, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %v", err)
	}

	//log.Printf("Content of file %s: %s\n", filePath, string(content))

	var data []Insight
	if err := json.Unmarshal(content, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	//fmt.Printf("Parsed data from file %s: %+v\n", filePath, data)
	return data, nil
}

type Insight struct {
	Country                string `json:"Country"`
	DeviceCategory         string `json:"DeviceCategory"`
	EngagementRate         string `json:"EngagementRate"`
	LandingPage            string `json:"LandingPage"`
	NewUsers               int    `json:"NewUsers"`
	ScreenPageViews        int    `json:"ScreenPageViews"`
	SessionMedium          string `json:"SessionMedium"`
	Sessions               int    `json:"Sessions"`
	TotalUsers             int    `json:"TotalUsers"`
	UserEngagementDuration int    `json:"UserEngagementDuration"`
	Date                   string `json:"date"`
}

func getInsightsFromLLM(apiKey string, data []Insight) (string, error) {
	prompt := createPrompt(data)
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-4o-mini",
		"messages": []RequestMessage{{
			Role:    "user",
			Content: prompt,
		},
		},
		"max_tokens": 1500,
	})
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	var aiResponse OpenAIResponse
	if err := json.Unmarshal(body, &aiResponse); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if len(aiResponse.Choices) == 0 {
		return "", fmt.Errorf("no choices in the response")
	}

	return aiResponse.Choices[0].Message.Content, nil
}

type RequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func createPrompt(data []Insight) string {
	prompt := "Analyze the following data and provide insights:\n\n"
	for i, d := range data {
		prompt += fmt.Sprintf("Country: %s, DeviceCategory: %s, EngagementRate: %s, LandingPage: %s, NewUsers: %d, ScreenPageViews: %d, SessionMedium: %s, Sessions: %d, TotalUsers: %d, UserEngagementDuration: %d, Date: %s\n",
			d.Country, d.DeviceCategory, d.EngagementRate, d.LandingPage, d.NewUsers, d.ScreenPageViews, d.SessionMedium, d.Sessions, d.TotalUsers, d.UserEngagementDuration, d.Date)
		if i == 100 {
			break
		}
	}
	//fmt.Println(prompt)
	//for i:= 0 ; i<100 ; i++ {
	//		prompt += fmt.Sprintf("Country: %s, DeviceCategory: %s, EngagementRate: %s, LandingPage: %s, NewUsers: %d, ScreenPageViews: %d, SessionMedium: %s, Sessions: %d, TotalUsers: %d, UserEngagementDuration: %d, Date: %s\n",
	//			d.Country, d.DeviceCategory, d.EngagementRate, d.LandingPage, d.NewUsers, d.ScreenPageViews, d.SessionMedium, d.Sessions, d.TotalUsers, d.UserEngagementDuration, d.Date)
	//	}
	//}
	return prompt
}

type OpenAIResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type Choice struct {
	Message      RequestMessage `json:"message"`
	Index        int            `json:"index"`
	Logprobs     interface{}    `json:"logprobs"`
	FinishReason string         `json:"finish_reason"`
}

type Choices []Choice

func (j Choices) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.Marshal(j)
}

func (j *Choices) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal HeartbeatModuleResult value:", value))
	}
	return json.Unmarshal(bytes, j)
}

func (j Choice) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Choice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal HeartbeatModuleResult value:", value))
	}
	return json.Unmarshal(bytes, j)
}

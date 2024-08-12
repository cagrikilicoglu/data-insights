package ai

import (
	"bytes"
	"data-insights/kit/model"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Client interface {
	makeRequestAndGetResponse(apiKey, prompt string) (OpenAIResponse, error)
	GetInsightsFromLLM(apiKey string, data model.UserMetrics) (string, error)
}

type OpenAIClient struct {
	AIModel    AIModel
	Url        string
	MaxTokens  int
	SenderRole string
}

func NewOpenAIClient(model AIModel, url string, maxTokens int, senderRole string) OpenAIClient {
	return OpenAIClient{
		AIModel:    model,
		Url:        url,
		MaxTokens:  maxTokens,
		SenderRole: senderRole,
	}
}

func (o OpenAIClient) makeRequestAndGetResponse(apiKey, prompt string) (OpenAIResponse, error) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": o.AIModel,
		"messages": []RequestMessage{{
			Role:    o.SenderRole,
			Content: prompt,
		},
		},
		"max_tokens": o.MaxTokens,
	})
	if err != nil {
		return OpenAIResponse{}, fmt.Errorf("failed to marshal request body: %v", err)
	}
	req, err := http.NewRequest("POST", o.Url, bytes.NewBuffer(requestBody))
	if err != nil {
		return OpenAIResponse{}, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return OpenAIResponse{}, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return OpenAIResponse{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var aiResponse OpenAIResponse
	if err := json.Unmarshal(body, &aiResponse); err != nil {
		return OpenAIResponse{}, fmt.Errorf("failed to unmarshal response: %v", err)
	}
	return aiResponse, nil
}

func (o OpenAIClient) GetInsightsFromLLM(apiKey string, data model.UserMetrics) (string, error) {

	aiResponse, err := o.makeRequestAndGetResponse(apiKey, createPrompt(data))
	if err != nil {
		return "", fmt.Errorf("failed to make request and get response: %v", err)
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

// todo delete??
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
		return errors.New(fmt.Sprint("Failed to unmarshal Choices value:", value))
	}
	return json.Unmarshal(bytes, j)
}

func (j Choice) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Choice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal Choice value:", value))
	}
	return json.Unmarshal(bytes, j)
}

package ai

import (
	"bytes"
	"data-insights/kit/common"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client interface {
	makeRequestAndGetResponse(apiKey, prompt string) (OpenAIResponse, error)
	GetInsightsFromLLM(apiKey string, data common.UserMetrics) (string, error)
}

type OpenAIClient struct {
	AIModel    AIModel
	Url        string
	MaxTokens  int
	SenderRole string
	HttpClient *http.Client // Reusable HTTP client
}

func NewOpenAIClient(model AIModel, url string, maxTokens int, senderRole string, httpClient *http.Client) OpenAIClient {
	return OpenAIClient{
		AIModel:    model,
		Url:        url,
		MaxTokens:  maxTokens,
		SenderRole: senderRole,
		HttpClient: httpClient,
	}
}

// makeRequestAndGetResponse creates an HTTP POST request to the OpenAI API with the provided API key and prompt,
// and returns the API response. It returns an OpenAIResponse object if successful, or an error if the request
// fails or the response cannot be parsed.

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

// GetInsightsFromLLM generates insights from the given UserMetrics data
// by sending a request to the OpenAI API and returning the first response choice
// as a string. It returns an error if the request fails or if the API returns no choices.
func (o OpenAIClient) GetInsightsFromLLM(apiKey string, data common.UserMetrics) (string, error) {

	aiResponse, err := o.makeRequestAndGetResponse(apiKey, createPrompt(data))
	if err != nil {
		return "", fmt.Errorf("failed to make request and get response: %v", err)
	}

	if len(aiResponse.Choices) == 0 {
		return "", fmt.Errorf("no choices in the response")
	}

	return aiResponse.Choices[0].Message.Content, nil
}

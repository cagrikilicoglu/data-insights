package ai

type AIModel string

const GPT4oMini AIModel = "gpt-4o-mini"

const (
	OpenAIUrl        = "https://api.openai.com/v1/chat/completions"
	OpenAIMaxTokens  = 1500
	OpenAISenderRole = "user"
)

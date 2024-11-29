package chatgpt

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

// ChatGPTClient estrutura para o cliente
type ChatGPTClient struct {
	client *openai.Client
}

// NewChatGPTClient inicializa o cliente ChatGPT
func NewChatGPTClient(apiKey string) *ChatGPTClient {
	return &ChatGPTClient{
		client: openai.NewClient(apiKey),
	}
}

// GetJarvisResponse envia uma mensagem e retorna a resposta simulando o Jarvis
func (c *ChatGPTClient) GetJarvisResponse(prompt string) (string, error) {
	ctx := context.Background()

	// Montar a mensagem

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are Jarvis, the AI assistant of Tony Stark. Respond as Jarvis would.",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}

	// Chamar a API
	resp, err := c.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo, // GPT-3.5-turbo
		Messages: messages,
	})
	if err != nil {
		return "", fmt.Errorf("error getting response from ChatGPT: %v", err)
	}

	// Retornar a resposta
	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response received from ChatGPT")
}

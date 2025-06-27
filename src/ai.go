package src

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

type Params struct {
	ApiKey        string
	BaseURL       string
	ModelName     string
	PromptSize    string
	MessageLength int
}

func getSystemPrompt(sysPrompt string) string {
	switch sysPrompt {
	case "mini":
		return SystemPromptMini
	case "big":
		return SystemPromptBig
	default:
		return SystemPromptMini
	}
}

func userPrompt(gitDiff string, answerLength int) string {
	return fmt.Sprintf(`
	Given the git changes below, please draft a concise commit message that accurately summarizes the modifications. Follow these guidelines:

	1. Limit your commit message to %d words.
	2. The whole commit message should in lowercase, no uppercase characters are allowed.
	3. Do not respond in Markdown

	   Git Changes: 

		`, answerLength) + gitDiff
}

func GenerateCommitMessage(gitDiff string, parms *Params) (string, error) {
	config := openai.DefaultConfig(parms.ApiKey)
	config.BaseURL = parms.BaseURL
	client := openai.NewClientWithConfig(config)

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: parms.ModelName,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: getSystemPrompt(parms.PromptSize),
			},
			{
				Role:    "user",
				Content: userPrompt(gitDiff, parms.MessageLength),
			},
		},
	})

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

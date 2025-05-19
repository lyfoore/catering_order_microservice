package service

import (
	"context"
	"github.com/openai/openai-go"
)

func GenerateChatCompletion(text string) (string, error) {
	client := openai.NewClient()
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(text),
		},
		Model: openai.ChatModelGPT4o,
	})
	if err != nil {
		return "Error occurred", err
	}

	return chatCompletion.Choices[0].Message.Content, nil
}

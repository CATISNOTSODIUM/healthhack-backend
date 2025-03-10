package api

import (
	"context"
	"log"
	openai "github.com/sashabaranov/go-openai" // unofficial library
	"github.com/sashabaranov/go-openai/jsonschema"
)

// This can be improved.
func ExtractTextFeature(text string, client * openai.Client) (string, error) {
	type Result struct {
		Coherence []struct {
			Score uint `json:"score"`
			Description string `json:"description"` 
		} `json:"coherence"`
		SentenceComplexity []struct {
			Score uint `json:"score"`
			Description string `json:"description"`
		} `json:"sentence_complexity"`
	}
	var result Result
	schema, err := jsonschema.GenerateSchemaForType(result)

	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "As a teacher, your goal is to evaluate the sentence (from speech) based on coherence and sentence complexity.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: text,
				},
			},
			ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
			JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
					Name:   "teacher",
					Schema: schema,
					Strict: true,
				},
			},
		},
	)

	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

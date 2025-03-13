package api

import (
	"context"
	"encoding/json"
	"log"

	openai "github.com/sashabaranov/go-openai" // unofficial library
	"github.com/sashabaranov/go-openai/jsonschema"
)

type Result struct {
	Coherence struct {
		Score uint `json:"score"`
		Description string `json:"description"` 
	} `json:"coherence"`
	SentenceComplexity struct {
		Score uint `json:"score"`
		Description string `json:"description"`
	} `json:"sentence_complexity"`
}

func ExtractTextFeature(text string, client * openai.Client) (*Result, error) {
	
	var result Result
	schema, err := jsonschema.GenerateSchemaForType(result)

	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return nil, err
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
		return nil, err
	}

	var returnResult Result
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &returnResult); err != nil {
        return nil, err // Means the string was invalid
    }

	return &returnResult, nil
}

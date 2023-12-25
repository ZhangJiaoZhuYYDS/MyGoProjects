package main

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
	"golang.org/x/net/context"
)

func main() {
	c := openai.NewClient("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")
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT3Ada,
		MaxTokens: 5,
		Prompt:    "Lorem ipsum",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Text)
}

package utilexamples

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

func StreamingRequest(client *genai.Client, ctx context.Context, model string, prompt string, config *genai.GenerateContentConfig) {
	content := genai.Text(prompt)
	responseStreaming := client.Models.GenerateContentStream(
		ctx,
		model,
		content,
		config,
	)

	for chunk := range responseStreaming {
		for _, part := range chunk.Candidates[0].Content.Parts {
			if len(part.Text) == 0 {
				continue
			}
			if part.Thought {
				fmt.Printf("Thought: %s\n", part.Text)
			} else {
				fmt.Printf("Answer: %s\n", part.Text)
			}
		}
	}
}

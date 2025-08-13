package utilexamples

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

func ChatGemini(ctx context.Context, client *genai.Client, history []*genai.Content, model string, config *genai.GenerateContentConfig, prompt string) (*genai.GenerateContentResponse, error) {
	chat, _ := client.Chats.Create(ctx, model, config, history)
	res, _ := chat.SendMessage(ctx, genai.Part{Text: prompt})

	JSONSimpleRequest(res)

	if len(res.Candidates) > 0 {
		fmt.Println(res.Candidates[0].Content.Parts[0].Text)
	}
	return res, nil
}

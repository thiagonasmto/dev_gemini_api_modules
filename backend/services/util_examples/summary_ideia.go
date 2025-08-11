package utilexamples

import (
	"fmt"

	"google.golang.org/genai"
)

func SummaryResponse(response *genai.GenerateContentResponse) {
	for _, part := range response.Candidates[0].Content.Parts {
		if part.Text != "" {
			if part.Thought {
				fmt.Println("Thoughts Summary:")
				fmt.Println(part.Text)
			} else {
				fmt.Println("Answer:")
				fmt.Println(part.Text)
			}
		}
	}
}

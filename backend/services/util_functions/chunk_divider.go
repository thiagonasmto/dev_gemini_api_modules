package utilfunctions

import "fmt"

type Chunk struct {
	ID       string
	Content  string
	Metadata map[string]string
}

func ChunkText(text, fileName string, chunkSize int) []Chunk {
	var chunks []Chunk
	runes := []rune(text)

	for i := 0; i < len(runes); i += chunkSize {
		end := i + chunkSize
		if end > len(runes) {
			end = len(runes)
		}
		chunks = append(chunks, Chunk{
			ID:      fmt.Sprintf("%s_%d", fileName, i/chunkSize),
			Content: string(runes[i:end]),
			Metadata: map[string]string{
				"file":  fileName,
				"index": fmt.Sprintf("%d", i/chunkSize),
			},
		})
	}
	return chunks
}

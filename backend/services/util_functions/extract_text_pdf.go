package utilfunctions

import (
	"os"
	"path/filepath"

	"github.com/ledongthuc/pdf"
)

func ExtractTextFromPDF(path string) (string, error) {
	// Descobre o diretório onde o binário está rodando
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	baseDir := filepath.Dir(exePath)

	// Junta o caminho base com o relativo passado
	absPath := filepath.Join(baseDir, path)

	f, r, err := pdf.Open(absPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var text string
	totalPage := r.NumPage()
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		page := r.Page(pageIndex)
		if page.V.IsNull() {
			continue
		}
		contents, _ := page.GetPlainText(nil)
		text += contents + "\n"
	}
	return text, nil
}

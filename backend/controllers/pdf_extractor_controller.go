package controllers

import (
	"fmt"
	utilfunctions "geminai_with_go/services/util_functions"

	"github.com/gin-gonic/gin"
)

func PdfExtractor(c *gin.Context) {
	response, err := utilfunctions.ExtractTextFromPDF("backend/docs/pdf_examples/webinar-ibef-sobena-apresentacao.pdf")
	if err != nil {
		fmt.Print("erro: ", err)
	}
	fmt.Println(response)
}

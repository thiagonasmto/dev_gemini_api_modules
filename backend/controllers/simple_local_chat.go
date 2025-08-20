package controllers

import (
	utilexamples "geminai_with_go/services/util_examples"

	"github.com/gin-gonic/gin"
)

func SimpleLocalChat(c *gin.Context) {
	utilexamples.LocalSimpleRequest()
}

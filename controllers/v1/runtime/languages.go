package runtime

import (
	"net/http"

	"github.com/gin-gonic/gin"

	runtimeService "rohandhamapurkar/code-executor/services/v1/runtime"
)

func GetSupportedLanguages(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"languages": runtimeService.PackagesJSON,
	})
}

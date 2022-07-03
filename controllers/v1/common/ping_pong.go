package common

import (
	"net/http"
	commonService "rohandhamapurkar/code-executor/services/v1/common"

	"github.com/gin-gonic/gin"
)

func PingPongHandler(ctx *gin.Context) {
	commonService.TestFunctionForPrinting()
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

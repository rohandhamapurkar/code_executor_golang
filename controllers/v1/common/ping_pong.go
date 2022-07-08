package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingPongHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

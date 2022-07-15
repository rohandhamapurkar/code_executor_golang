package v1

import (
	commonController "rohandhamapurkar/code-executor/controllers/v1/common"

	"github.com/gin-gonic/gin"
)

func SetCommonControllerRoutes(rg *gin.RouterGroup) {
	commonGroup := rg.Group("/common")
	commonGroup.GET("/ping", commonController.PingPongHandler)
}

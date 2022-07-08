package v1

import (
	commonController "rohandhamapurkar/code-executor/controllers/v1/common"
	authMiddleware "rohandhamapurkar/code-executor/middlewares/auth"
	middleware "rohandhamapurkar/code-executor/middlewares/custom"

	"github.com/gin-gonic/gin"
)

func SetCommonControllerRoutes(rg *gin.RouterGroup) {
	commonGroup := rg.Group("/common")
	commonGroup.GET("/ping", middleware.CustomMiddleware(), authMiddleware.IsLoggedIn(), commonController.PingPongHandler)
}

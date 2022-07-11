package v1

import (
	runtimeController "rohandhamapurkar/code-executor/controllers/v1/runtime"
	authMiddleware "rohandhamapurkar/code-executor/middlewares/auth"

	"github.com/gin-gonic/gin"
)

func SetRuntimeControllerRoutes(rg *gin.RouterGroup) {
	commonGroup := rg.Group("/runtime")
	commonGroup.POST("/execute", authMiddleware.IsLoggedIn(), runtimeController.ExecuteCodeHandler)
}

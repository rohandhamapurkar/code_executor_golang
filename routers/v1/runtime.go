package v1

import (
	runtimeController "rohandhamapurkar/code-executor/controllers/v1/runtime"
	httpMiddlewares "rohandhamapurkar/code-executor/middlewares/httpmiddlewares"

	"github.com/gin-gonic/gin"
)

func SetRuntimeControllerRoutes(rg *gin.RouterGroup) {
	runtimeGroup := rg.Group("/runtime")
	runtimeGroup.POST("/execute", httpMiddlewares.RateLimit(), runtimeController.ExecuteCodeHandler)
	runtimeGroup.GET("/supported-languages", runtimeController.GetSupportedLanguages)
}

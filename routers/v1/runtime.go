package v1

import (
	runtimeController "rohandhamapurkar/code-executor/controllers/v1/runtime"
	httpMiddleware "rohandhamapurkar/code-executor/middlewares/http"

	"github.com/gin-gonic/gin"
)

func SetRuntimeControllerRoutes(rg *gin.RouterGroup) {
	runtimeGroup := rg.Group("/runtime")
	runtimeGroup.POST("/execute", httpMiddleware.RateLimit(), runtimeController.ExecuteCodeHandler)
	runtimeGroup.GET("/supported-languages", runtimeController.GetSupportedLanguages)
}

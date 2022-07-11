package v1

import "github.com/gin-gonic/gin"

func SetV1Routes(router *gin.Engine) {
	v1RouterGroup := router.Group("/v1")
	SetCommonControllerRoutes(v1RouterGroup)
	SetRuntimeControllerRoutes(v1RouterGroup)
}

package v1

import (
	authController "rohandhamapurkar/code-executor/controllers/v1/auth"

	"github.com/gin-gonic/gin"
)

func SetAuthControllerRoutes(rg *gin.RouterGroup) {
	authGroup := rg.Group("/auth")
	authGroup.POST("/register", authController.RegisterUserHandler)
}

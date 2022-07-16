package v1

import (
	snippetsController "rohandhamapurkar/code-executor/controllers/v1/snippets"
	middleware "rohandhamapurkar/code-executor/middlewares/auth"

	"github.com/gin-gonic/gin"
)

func SetSnippetsControllerRoutes(rg *gin.RouterGroup) {
	snippetsGroup := rg.Group("/snippets")
	snippetsGroup.POST("/", middleware.IsLoggedIn(), snippetsController.CreateSnippetHandler)
	snippetsGroup.GET("/", middleware.IsLoggedIn(), snippetsController.ReadSnippetHandler)
	snippetsGroup.PATCH("/", middleware.IsLoggedIn(), snippetsController.UpdateSnippetHandler)
	snippetsGroup.DELETE("/", middleware.IsLoggedIn(), snippetsController.DeleteSnippetHandler)
}

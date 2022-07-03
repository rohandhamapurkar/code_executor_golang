package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// custom middleware for consoling
func CustomMiddleware() gin.HandlerFunc {
	// Perform initialization here...
	return func(ctx *gin.Context) {
		fmt.Println("I'm a middleware :)")
		ctx.Next()
	}
}

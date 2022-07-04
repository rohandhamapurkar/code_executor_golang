package common

import (
	"log"

	"github.com/gin-gonic/gin"
)

// custom middleware for consoling
func CustomMiddleware() gin.HandlerFunc {
	// Perform initialization here...
	return func(ctx *gin.Context) {
		log.Println("I'm a middleware :)")
		ctx.Next()
	}
}

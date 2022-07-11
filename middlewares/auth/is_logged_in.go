package auth

import (
	"log"
	"net/http"
	"rohandhamapurkar/code-executor/core/constants"
	"rohandhamapurkar/code-executor/services/v1/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsLoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// extract token string from bearer token
		authorization := strings.Replace(ctx.GetHeader("Authorization"), "Bearer ", "", 1)

		// if no token string found then
		if authorization == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": constants.INVALID_TOKEN, "message": constants.MISSING_AUTH_TOKEN})
			return
		}

		// decode and verify the jwt token
		jwtClaims, err := auth.DecodeAndVerifyJwtToken(authorization)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": constants.INVALID_TOKEN, "message": constants.INVALID_TOKEN})
			return
		}

		ctx.Set("jwt_claims", jwtClaims)

		ctx.Next()
	}
}

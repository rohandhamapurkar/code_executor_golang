package auth

import (
	"net/http"
	"rohandhamapurkar/code-executor/core/constants"
	"rohandhamapurkar/code-executor/core/structs"
	"rohandhamapurkar/code-executor/services/v1/auth"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func IsLoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// extract token string from bearer token
		authorization := strings.Replace(ctx.GetHeader("Authorization"), "Bearer ", "", 1)

		// if no token string found then
		if authorization == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": constants.INVALID_TOKEN, "message": constants.MISSING_AUTH_TOKEN})
			ctx.Abort()
			return
		}

		// decode and verify the jwt token
		jwtToken, err := auth.DecodeAndVerifyJwtToken(authorization)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": constants.INVALID_TOKEN, "message": constants.INVALID_TOKEN})
			ctx.Abort()
			return
		}

		err = jwtToken.Claims.Valid()
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": constants.INVALID_TOKEN, "message": constants.INVALID_TOKEN})
			ctx.Abort()
			return
		}

		ctx.Set("jwt_claims", &structs.JWTClaims{
			AuthTime: jwtToken.Claims.(jwt.MapClaims)["auth_time"].(float64),
			ClientId: jwtToken.Claims.(jwt.MapClaims)["client_id"].(string),
			Exp:      jwtToken.Claims.(jwt.MapClaims)["exp"].(float64),
			Iat:      jwtToken.Claims.(jwt.MapClaims)["iat"].(float64),
			Iss:      jwtToken.Claims.(jwt.MapClaims)["iss"].(string),
			Jti:      jwtToken.Claims.(jwt.MapClaims)["jti"].(string),
			Scope:    jwtToken.Claims.(jwt.MapClaims)["scope"].(string),
			Sub:      jwtToken.Claims.(jwt.MapClaims)["sub"].(string),
			TokenUse: jwtToken.Claims.(jwt.MapClaims)["token_use"].(string),
			Username: jwtToken.Claims.(jwt.MapClaims)["username"].(string),
			Version:  jwtToken.Claims.(jwt.MapClaims)["version"].(float64),
		})

		ctx.Next()
	}
}

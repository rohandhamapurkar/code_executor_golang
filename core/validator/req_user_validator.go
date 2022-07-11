package validator

import (
	"net/http"
	"rohandhamapurkar/code-executor/core/constants"
	"rohandhamapurkar/code-executor/core/structs"

	"github.com/gin-gonic/gin"
)

func GetRequestUserClaims(ctx *gin.Context) *structs.JWTClaims {
	value, exists := ctx.Get("jwt_claims")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{
				"error":   constants.ERROR_MISSING_JWT_CLAIMS,
				"message": constants.MISSING_JWT_CLAIMS,
			})
		return nil
	}

	claims, ok := value.(*structs.JWTClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{
				"error":   constants.ERROR_MISSING_JWT_CLAIMS,
				"message": constants.MISSING_JWT_CLAIMS,
			})
		return nil
	}

	return claims
}

package validator

import (
	"net/http"
	"rohandhamapurkar/code-executor/core/constants"

	"github.com/gin-gonic/gin"
)

func ParseAndValidateRequestBody(ctx *gin.Context, requestBody interface{}) bool {
	errEncountered := false
	if err := ctx.ShouldBindJSON(requestBody); err != nil {
		errEncountered = true
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   constants.ERROR_INVALID_REQUEST_BODY,
				"message": err.Error(),
			})

	}

	errors := validateStruct(requestBody)
	if len(errors) > 0 {
		errEncountered = true
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   constants.ERROR_INVALID_REQUEST_BODY,
				"message": errors,
			})
	}
	return errEncountered
}

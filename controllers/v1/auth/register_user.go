package auth

import (
	"log"
	"net/http"

	"rohandhamapurkar/code-executor/core/constants"
	"rohandhamapurkar/code-executor/core/structs"
	validator "rohandhamapurkar/code-executor/core/validator"
	authService "rohandhamapurkar/code-executor/services/v1/auth"

	"github.com/gin-gonic/gin"
)

func RegisterUserHandler(ctx *gin.Context) {
	user := &structs.UserRegReqBody{}
	if err := ctx.ShouldBindJSON(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   constants.ERROR_INVALID_REQUEST_BODY,
				"message": err.Error(),
			})
		return
	}

	errors := validator.ValidateStruct(user)
	if len(errors) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   constants.ERROR_INVALID_REQUEST_BODY,
				"message": errors,
			})
		return
	}

	if err := authService.RegisterUser(user); err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   constants.USER_REG_FAILED,
				"message": constants.ERROR_USER_REG_SERVICE_UNAVAILABLE,
			})
		return
	}

	ctx.JSON(http.StatusCreated, &gin.H{
		"message": constants.USER_REG_SUCCESS,
	})
}

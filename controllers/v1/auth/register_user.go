package auth

import (
	"log"
	"net/http"

	"rohandhamapurkar/code-executor/core/constants"
	authService "rohandhamapurkar/code-executor/services/v1/auth"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type RequestBody struct {
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=8,max=32,alphanum"`
	ConfirmPassword string `json:"confirmPassword" binding:"eqfield=Password,required"`
}

func RegisterUserHandler(ctx *gin.Context) {
	body := RequestBody{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   constants.ERROR_INVALID_REQUEST_BODY,
				"message": err.Error(),
			})
		return
	}

	if err := authService.RegisterUser(&body.Email, &body.Password); err != nil {
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

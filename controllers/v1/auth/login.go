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

func LoginHandler(ctx *gin.Context) {
	body := &structs.LoginReqBody{}
	if err := validator.ParseAndValidateRequestBody(ctx, body); err == true {
		return
	}

	jwtString, err := authService.LoginUser(body)

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   constants.USER_LOGIN_FAILED,
				"message": err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusAccepted, &gin.H{
		"message":      constants.USER_LOGIN_SUCCESS,
		"access_token": jwtString,
	})
}

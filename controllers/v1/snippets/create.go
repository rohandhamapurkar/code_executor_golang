package snippets

import (
	"net/http"
	"rohandhamapurkar/code-executor/core/constants"
	"rohandhamapurkar/code-executor/core/structs"
	"rohandhamapurkar/code-executor/core/validator"
	snippetsService "rohandhamapurkar/code-executor/services/v1/snippets"

	"github.com/gin-gonic/gin"
)

func CreateSnippetHandler(ctx *gin.Context) {
	body := &structs.CreateSnippetReqBody{}
	if err := validator.ParseAndValidateRequestBody(ctx, body); err {
		return
	}

	var userClaims *structs.JWTClaims
	if userClaims = validator.GetRequestUserClaims(ctx); userClaims == nil {
		return
	}

	insertedID, err := snippetsService.CreateSnippet(body, userClaims.Sub)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": constants.ERROR_SNIPPET_CREATION, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": constants.SNIPPET_CREATED, "insertedID": insertedID})

}

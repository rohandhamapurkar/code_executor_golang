package snippets

import (
	"net/http"
	"rohandhamapurkar/code-executor/core/constants"
	"rohandhamapurkar/code-executor/core/structs"
	"rohandhamapurkar/code-executor/core/validator"
	snippetsService "rohandhamapurkar/code-executor/services/v1/snippets"

	"github.com/gin-gonic/gin"
)

func DeleteSnippetHandler(ctx *gin.Context) {
	body := &structs.DeleteSnippetReqBody{}
	if err := validator.ParseAndValidateRequestBody(ctx, body); err {
		return
	}

	var userClaims *structs.JWTClaims
	if userClaims = validator.GetRequestUserClaims(ctx); userClaims == nil {
		return
	}

	err := snippetsService.DeleteSnippet(body, userClaims.Sub)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": constants.ERROR_SNIPPET_DELETION, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": constants.SNIPPET_DELETED})
}

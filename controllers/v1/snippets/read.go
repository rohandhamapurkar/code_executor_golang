package snippets

import (
	"net/http"
	"rohandhamapurkar/code-executor/core/constants"
	"rohandhamapurkar/code-executor/core/structs"
	"rohandhamapurkar/code-executor/core/validator"

	snippetsService "rohandhamapurkar/code-executor/services/v1/snippets"

	"github.com/gin-gonic/gin"
)

func ReadSnippetHandler(ctx *gin.Context) {

	var userClaims *structs.JWTClaims
	if userClaims = validator.GetRequestUserClaims(ctx); userClaims == nil {
		return
	}

	snippets, err := snippetsService.GetSnippet(userClaims.Sub)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": constants.ERROR_SNIPPET_READ, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": snippets})

}

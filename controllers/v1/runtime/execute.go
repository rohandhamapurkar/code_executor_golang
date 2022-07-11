package runtime

import (
	"log"
	"net/http"
	"os"
	"rohandhamapurkar/code-executor/core/structs"
	"rohandhamapurkar/code-executor/core/validator"

	"github.com/gin-gonic/gin"
)

func ExecuteCodeHandler(ctx *gin.Context) {
	body := &structs.ExecuteCodeReqBody{}
	if err := validator.ParseAndValidateRequestBody(ctx, body); err {
		return
	}
	log.Println(body)

	var userClaims *structs.JWTClaims
	if userClaims = validator.GetRequestUserClaims(ctx); userClaims == nil {
		return
	}
	log.Println(os.TempDir())

	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

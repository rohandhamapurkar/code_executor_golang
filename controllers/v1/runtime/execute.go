package runtime

import (
	"log"
	"net/http"
	"rohandhamapurkar/code-executor/core/constants"
	"rohandhamapurkar/code-executor/core/structs"
	"rohandhamapurkar/code-executor/core/validator"
	runtimeService "rohandhamapurkar/code-executor/services/v1/runtime"

	"github.com/gin-gonic/gin"
)

func ExecuteCodeHandler(ctx *gin.Context) {
	body := &structs.ExecuteCodeReqBody{}
	if err := validator.ParseAndValidateRequestBody(ctx, body); err {
		return
	}

	output, err := runtimeService.SafeCallLibrary(body)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error":   constants.ERROR_EXECUTION_ERROR,
				"message": err.Error(),
				"stdout":  output.StdOut,
				"stderr":  output.StdErr,
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"stdout": output.StdOut, "stderr": output.StdErr})
}

package validator

import (
	"rohandhamapurkar/code-executor/services/v1/runtime"

	"github.com/go-playground/validator/v10"
)

func isProgrammingLanguageSupported(fl validator.FieldLevel) bool {
	language := fl.Field().String()
	for key := range runtime.Packages {
		if key == language {
			return true
		}
	}
	return false
}

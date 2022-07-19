package validator

import (
	"log"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	defer log.Println("Initialized validator")
	validate = validator.New()
	validate.RegisterValidation("isProgrammingLanguageSupported", isProgrammingLanguageSupported)
}

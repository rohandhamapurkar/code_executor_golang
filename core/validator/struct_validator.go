package validator

import (
	"rohandhamapurkar/code-executor/core/structs"

	"github.com/go-playground/validator/v10"
)

func validateStruct(dto interface{}) []*structs.ErrorResponse {
	validate := validator.New()
	var errors []*structs.ErrorResponse
	err := validate.Struct(dto)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element structs.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

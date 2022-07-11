package structs

type ExecuteCodeReqBody struct {
	Language string `validate:"required"`
	Code     string `validate:"required"`
}

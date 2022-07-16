package structs

type ExecuteCodeReqBody struct {
	Language string `json:"language" validate:"required" `
	Code     string `json:"code" validate:"required"`
}

type CreateSnippetReqBody struct {
	Name     string `json:"name" validate:"required"`
	Language string `json:"language" validate:"required"`
	Code     string `json:"code" validate:"required"`
	Public   *bool  `json:"makePublic" validate:"required"`
}

type UpdateSnippetReqBody struct {
	ID       uint   `json:"ID" validate:"required"`
	Name     string `json:"name" validate:"omitempty"`
	Language string `json:"language" validate:"omitempty"`
	Code     string `json:"code" validate:"omitempty"`
	Public   *bool  `json:"makePublic" validate:"omitempty"`
}

type DeleteSnippetReqBody struct {
	ID uint `json:"ID" validate:"required"`
}

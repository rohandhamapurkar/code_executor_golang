package structs

type ErrorResponse struct {
	FailedField string      `json:"failedField"`
	Tag         string      `json:"tag"`
	ParamRecv   string      `json:"paramRecv"`
	ValueRecv   interface{} `json:"valueRecv"`
}

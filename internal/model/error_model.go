package model

type ErrorDetail struct {
	Field     string `json:"field"`
	ErrorCode string `json:"error_code"`
	Param     string `json:"param"`
	Message   string `json:"message"`
}

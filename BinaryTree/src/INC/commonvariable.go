package INC

type ErrorResponse struct {
	Error_code    int    `json:"error_code"`
	Error_message string `json:"error_message"`
}
var Error_details ErrorResponse


package response

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

type Response struct {
	StatusCode   int    `json:"status"`
	Message string `json:"message"`
}

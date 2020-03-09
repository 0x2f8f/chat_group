package response

import (
	"net/http"
	"log"
	"encoding/json"
)

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func GetError(err error, w http.ResponseWriter, statusCode int, errorResponseMessage string) {
	log.Println(errorResponseMessage)
	log.Println(err.Error())
	var response = ErrorResponse{
		ErrorMessage: errorResponseMessage,
		StatusCode:   statusCode,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
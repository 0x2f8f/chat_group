package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func MessageCreateHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)

	var requestData struct {
		ChatId string `json:"chat_id"`
		Message string `json:"message"`
	}

	err := decoder.Decode(&requestData)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Create message by chat: %v\n Message: %v", requestData.ChatId, requestData.Message)
}

func MessageEditHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Message edit: %v\nNew message: %v", vars["message_id"], vars["message"])
}

func MessageDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Message delete: %v\n", vars["message_id"])
}

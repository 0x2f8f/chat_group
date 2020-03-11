package handlers

import (
	"net/http"
	"fmt"
	"chat_group/response"
	"log"
	"encoding/json"
)

type AuthPhoneRequest struct {
	Uid string `json:"uid"`
	Phone string `json:"phone"`
	Phone2 string `json:"phone2"`
}

func AuthPhoneHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req AuthPhoneRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	if (len(req.Uid) == 0) {
		response.GetErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Param not found: %v", "uid"))

		return
	}

	if (len(req.Phone) == 0) {
		response.GetErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Param not found: %v", "phone"))

		return
	}

	log.Printf("Login attempt: uid - %v, phone - %v", req.Uid, req.Phone)

	var response = response.Response{
		StatusCode: http.StatusOK,
		Message: fmt.Sprint("Need confirm"),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func AuthConfirmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
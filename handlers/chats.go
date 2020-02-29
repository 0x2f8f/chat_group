package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GroupCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Create chat group")
}

func GroupInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Info by chat: %v\n", vars["chat_id"])
}

func GroupEditHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Chat edit: %v\n", vars["chat_id"])
}

func GroupDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Chat delete: %v\n", vars["chat_id"])
}

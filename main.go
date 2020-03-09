package main

import (
	"chat_group/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func main() {
	r := mux.NewRouter()

	g := r.PathPrefix("/chat/group").Subrouter()
	g.HandleFunc("", handlers.GroupCreateHandler).Methods("POST")
	g.HandleFunc("/{chat_id}", handlers.GroupInfoHandler).Methods("GET")
	g.HandleFunc("/{chat_id}", handlers.GroupEditHandler).Methods("PUT")
	g.HandleFunc("/{chat_id}", handlers.GroupDeleteHandler).Methods("DELETE")

	m := r.PathPrefix("/chat/message").Subrouter()
	m.HandleFunc("", handlers.MessageCreateHandler).Methods("POST")
	m.HandleFunc("/{message_id}", handlers.MessageEditHandler).Methods("PUT")
	m.HandleFunc("/{message_id}", handlers.MessageDeleteHandler).Methods("DELETE")

	http.ListenAndServe(":667", r)
	fmt.Println("Start Chat App on 667 port")
}

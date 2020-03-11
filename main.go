package main

import (
	"chat_group/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"log"
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

	a := r.PathPrefix("/auth").Subrouter()
	a.HandleFunc("/phone", handlers.AuthPhoneHandler).Methods("POST")
	a.HandleFunc("/confirm", handlers.AuthConfirmHandler).Methods("POST")

	http.ListenAndServe(":678", r)
	log.Println("Start Chat App on 678 port")
}

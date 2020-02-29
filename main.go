package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create chat group")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/chat-group", createHandler).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":666", nil)

}
package main

import (
	"fmt"
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create chat group")
}

func main() {
	fmt.Println("http://localhost:666/")

	http.HandleFunc("/chat-group", createHandler)
	//http.HandleFunc("/SavePost", savePostHandler)

	http.ListenAndServe(":666", nil)

}
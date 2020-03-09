package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"context"
	"encoding/json"
	"chat_group/models"
	"chat_group/mongodb"
)

func GroupCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var group models.Group
	_ = json.NewDecoder(r.Body).Decode(&group)

	collection := mongodb.GetCollection("croups")
	insertResult, err := collection.InsertOne(context.TODO(), group)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertResult)
	//fmt.Fprintf(w, "Create chat group %s", insertResult.InsertedID)
}

func GroupInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Create chat group!")
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

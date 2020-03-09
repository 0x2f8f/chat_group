package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"context"
	"encoding/json"
	"chat_group/models"
	"chat_group/mongodb"
	"chat_group/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

func GroupCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var group models.Group
	_ = json.NewDecoder(r.Body).Decode(&group)

	collection := mongodb.GetCollection("croups")
	insertResult, err := collection.InsertOne(context.TODO(), group)
	if err != nil {
		response.GetError(err, w, http.StatusInternalServerError, err.Error())

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertResult)
}

func GroupInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var group models.Group
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["chat_id"])

	collection := mongodb.GetCollection("croups")
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&group)

	if err != nil {
		response.GetError(err, w, http.StatusNotFound, fmt.Sprintf("Chat not found: %v", params["chat_id"]))

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(group)
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

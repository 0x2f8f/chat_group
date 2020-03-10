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

	err = collection.FindOne(context.TODO(), bson.M{"_id": insertResult.InsertedID}).Decode(&group)
	if err != nil {
		response.GetError(err, w, http.StatusNotFound, fmt.Sprintf("failed to create chat"))

		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(group)
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
	w.Header().Set("Content-Type", "application/json")

	var group models.Group
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["chat_id"])

	collection := mongodb.GetCollection("croups")
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&group)

	if err != nil {
		response.GetError(err, w, http.StatusNotFound, fmt.Sprintf("Chat not found: %v", params["chat_id"]))

		return
	}

	_ = json.NewDecoder(r.Body).Decode(&group)
	update := bson.D{
		{"$set", bson.D{
			{"name", group.Name},
			{"description", group.Description},
		}},
	}

	err = collection.FindOneAndUpdate(context.TODO(), bson.M{"_id": id}, update).Decode(&group)
	if err != nil {
		response.GetError(err, w, http.StatusNotFound, fmt.Sprintf("Failed to update chat %s", params["chat_id"]))

		return
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&group)
	if err != nil {
		response.GetError(err, w, http.StatusNotFound, fmt.Sprintf("Failed to update chat %s", params["chat_id"]))

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(group)
}

func GroupDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var group models.Group
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["chat_id"])

	collection := mongodb.GetCollection("croups")
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&group)
	if err != nil {
		response.GetError(err, w, http.StatusNotFound, fmt.Sprintf("Chat not found: %v", params["chat_id"]))

		return
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		response.GetError(err, w, http.StatusInternalServerError, err.Error())

		return
	}

	var response = response.Response{
		StatusCode: http.StatusOK,
		Message: fmt.Sprintf("Chat deleted successfully: %v", params["chat_id"]),

	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

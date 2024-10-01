package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-list/database"
	"todo-list/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateChecklist godoc
// @Summary Create a new checklist
// @Description Create a new checklist for the logged-in user
// @Tags Checklists
// @Accept  json
// @Produce  json
// @Param checklist body models.Checklist true "Checklist info"
// @Success 201 {object} models.Checklist
// @Failure 400 {string} string "Invalid input"
// @Router /checklists [post]
func CreateChecklist(w http.ResponseWriter, r *http.Request) {
	var checklist models.Checklist
	err := json.NewDecoder(r.Body).Decode(&checklist)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	collection := database.DB.Collection("checklists")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, checklist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(checklist)
}

// GetChecklists godoc
// @Summary Get all checklists
// @Description Get all checklists for the logged-in user
// @Tags Checklists
// @Produce  json
// @Success 200 {array} models.Checklist
// @Router /checklists [get]
func GetChecklists(w http.ResponseWriter, r *http.Request) {
	collection := database.DB.Collection("checklists")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Could not retrieve checklists", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var checklists []models.Checklist
	for cursor.Next(ctx) {
		var checklist models.Checklist
		cursor.Decode(&checklist)
		checklists = append(checklists, checklist)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(checklists)
}

// DeleteChecklist godoc
// @Summary Delete a checklist
// @Description Delete a checklist by ID
// @Tags Checklists
// @Param id path string true "Checklist ID"
// @Success 204 {string} string "Checklist deleted successfully"
// @Failure 404 {string} string "Checklist not found"
// @Failure 500 {string} string "Failed to delete checklist"
// @Router /checklists/{id} [delete]
func DeleteChecklist(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid checklist ID", http.StatusBadRequest)
		return
	}

	collection := database.DB.Collection("checklists")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		http.Error(w, "Failed to delete checklist", http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "Checklist not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

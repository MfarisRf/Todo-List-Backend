package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"todo-list/database"
	"todo-list/middlewares"
	"todo-list/models"

	"go.mongodb.org/mongo-driver/bson"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with username and password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body models.User true "User info"
// @Success 201 {object} models.User
// @Failure 400 {string} string "Invalid input"
// @Router /register [post]
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	collection := database.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// LoginUser godoc
// @Summary Login a user
// @Description Login a user and return JWT token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body models.User true "User info"
// @Success 200 {object} string
// @Failure 401 {string} string "Invalid credentials"
// @Router /login [post]
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	collection := database.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var foundUser models.User
	err = collection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundUser)
	if err != nil || foundUser.Password != user.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	tokenString, err := middlewares.GenerateJWT(foundUser.ID.Hex())
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

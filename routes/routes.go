package routes

import (
	"net/http"
	"todo-list/controllers"
	"todo-list/middlewares"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// Rute untuk autentikasi
	// @Summary Register a new user
	// @Description Register a new user
	// @Tags Auth
	// @Accept json
	// @Produce json
	// @Param user body models.User true "User info"
	// @Success 201 {object} models.User
	// @Failure 400 {string} string "Invalid input"
	// @Router /register [post]
	r.HandleFunc("/register", controllers.RegisterUser).Methods("POST")

	// @Summary Login user
	// @Description Login user to get a JWT
	// @Tags Auth
	// @Accept json
	// @Produce json
	// @Param user body models.User true "User info"
	// @Success 200 {string} string "JWT token"
	// @Failure 400 {string} string "Invalid input"
	// @Router /login [post]
	r.HandleFunc("/login", controllers.LoginUser).Methods("POST")

	r.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger"))))

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.JwtVerify)

	// Rute untuk checklist
	// @Summary Create a new checklist
	// @Description Create a new checklist for the logged-in user
	// @Tags Checklists
	// @Accept json
	// @Produce json
	// @Param checklist body models.Checklist true "Checklist info"
	// @Success 201 {object} models.Checklist
	// @Failure 400 {string} string "Invalid input"
	// @Router /checklists [post]
	api.HandleFunc("/checklists", controllers.CreateChecklist).Methods("POST")

	// @Summary Get all checklists
	// @Description Get all checklists for the logged-in user
	// @Tags Checklists
	// @Produce json
	// @Success 200 {array} models.Checklist
	// @Router /checklists [get]
	api.HandleFunc("/checklists", controllers.GetChecklists).Methods("GET")

	// @Summary Delete a checklist
	// @Description Delete a checklist by ID
	// @Tags Checklists
	// @Accept json
	// @Produce json
	// @Param id path string true "Checklist ID"
	// @Success 204 {string} string "No Content"
	// @Failure 404 {string} string "Checklist not found"
	// @Router /checklists/{id} [delete]
	api.HandleFunc("/checklists/{id}", controllers.DeleteChecklist).Methods("DELETE")

	return r
}

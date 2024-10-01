package main

import (
	"log"
	"net/http"
	"todo-list/database"
	"todo-list/routes"
)

func main() {

	database.ConnectDB()
	defer database.DisconnectDB()

	r := routes.RegisterRoutes()

	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

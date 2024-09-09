package main

import (
	"log"
	"net/http"
	"github.com/gopsdv/lightweight/handlers"
	"github.com/gopsdv/lightweight/database"
)

func main() {
	// Initialize the database connection
	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Ensure the database connection is closed when the program exits
	defer database.CloseDB()

	log.Println("Starting server at port 8080")
	http.HandleFunc("/exercises", handlers.ExercisesHandler)
	http.ListenAndServe(":8080", nil)
}
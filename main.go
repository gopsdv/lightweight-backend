package main

import (
	"github.com/gopsdv/lightweight/database"
	"github.com/gopsdv/lightweight/handlers"
	"log"
	"net/http"
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

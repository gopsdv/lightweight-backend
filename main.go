package main

import (
	"github.com/gopsdv/lightweight/database"
	"github.com/gopsdv/lightweight/handlers/exercises"
	"github.com/gopsdv/lightweight/handlers/sets"
	"github.com/gopsdv/lightweight/handlers/workouts"
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
	http.HandleFunc("/sets", sets.MethodHandler)
	http.HandleFunc("/exercises", exercises.MethodHandler)
	http.HandleFunc("/workouts", workouts.MethodHandler)
	http.ListenAndServe(":8080", nil)
}

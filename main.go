package main

import (
	"fmt"
	"log"
	"lightweight/handlers"
	"github.com/joho/godotenv"
)

type Set struct {
	Weight float32
	Reps uint8
	PartialReps uint8
	RIR uint8 // Reps in Reserve
}

type Exercise struct { 
	ID int `json:"id"`
	Name string `json:"name"`
}



func main() {
	env, _ := godotenv.Read(".env")
	var port string = env["PORT"]
	if port == "" {
		log.Fatal("PORT not found")
	}
	fmt.Println("Starting server at port 8080")
	http.HandleFunc("/bro", exerciseHandler)
	http.ListenAndServe(":8080", nil)
}
package sets

import (
	"encoding/json"
	"log"
	"net/http"
)


type Exercise struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func POST(w http.ResponseWriter, r *http.Request) {

	var exercise Exercise

	if err := json.NewDecoder(r.Body).Decode(&exercise); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	ExerciseID, err := AddExercise(exercise)
	if err != nil {
		log.Fatalln(err)
	}

	response := map[string]interface{}{"error": false, "message": "Created a new exercise", "data": map[string]interface{}{"exerciseID": ExerciseID}}
	w.Header().Set("Content-Type", "application/json")
	resErr := json.NewEncoder(w).Encode(response)
	if resErr != nil {
		log.Print("Request Successfull")
	}
}

func GET(w http.ResponseWriter, r *http.Request) {

	exersices, err := getExercises()
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Failed to fetch excersices", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exersices)
}

func ExercisesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		POST(w, r)
	case http.MethodGet:
		GET(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

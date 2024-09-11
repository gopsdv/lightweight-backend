package workouts

import (
	"encoding/json"
	"log"
	"net/http"
)

type Workout struct {
	ID         uint16
	ExerciseID uint8
	Date       string
}

func POST(w http.ResponseWriter, r *http.Request) {

	var workout Workout

	if err := json.NewDecoder(r.Body).Decode(&workout); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	WorkoutID, err := AddWorkout(workout)
	if err != nil {
		log.Fatalln(err)
	}

	response := map[string]interface{}{"error": false, "message": "A Workout added", "data": map[string]interface{}{"workoutID": WorkoutID}}

	w.Header().Set("Content-Type", "application/json")
	resErr := json.NewEncoder(w).Encode(response)
	if resErr != nil {
		log.Print("Request Successfull")
	}
}

func GET(w http.ResponseWriter, r *http.Request) {

	workouts, err := getWorkouts()
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Failed to fetch excersices", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workouts)
}

func MethodHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		POST(w, r)
	case http.MethodGet:
		GET(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

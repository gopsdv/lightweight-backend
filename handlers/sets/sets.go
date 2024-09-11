package sets

import (
	"encoding/json"
	"log"
	"net/http"
)

type Set struct {
	ID          uint32
	WorkoutID   uint8
	SetNum      uint8
	Weight      float32
	Reps        uint8
	PartialReps uint8
	RIR         uint8 // Reps in Reserve
}

func POST(w http.ResponseWriter, r *http.Request) {

	var workingSet Set

	if err := json.NewDecoder(r.Body).Decode(&workingSet); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	setID, err := AddSet(workingSet)
	if err != nil {
		log.Fatalln(err)
	}

	response := map[string]interface{}{"error": false, "message": "Added a working set", "data": map[string]interface{}{"setID": setID}}
	w.Header().Set("Content-Type", "application/json")
	resErr := json.NewEncoder(w).Encode(response)
	if resErr != nil {
		log.Print("Request Successfull")
	}
}

func GET(w http.ResponseWriter, r *http.Request) {

	exersices, err := getSets()
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Failed to fetch excersices", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exersices)
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

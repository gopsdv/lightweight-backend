package handlers

import (
	"fmt"
	"testing"
	"github.com/gopsdv/lightweight/database"
)

func TestAddExercise(t *testing.T) {
	database.Connect()
	defer database.CloseDB()
	ExerciseID, _ := AddExercise(database.Exercise{Name: "Barbell Back Squat"})
	fmt.Println(ExerciseID)
}


func TestGetExercise(t *testing.T) {
	database.Connect()
	defer database.CloseDB()
	rows , err := getExercises()
	if err != nil {
		fmt.Println("Rows: ", err)
	}
	fmt.Println("Rows: ", rows)

}
package workouts

import (
	"fmt"
	"github.com/gopsdv/lightweight/database"
)

func AddWorkout(workout Workout) (int64, error) {
	result, err := database.DB.Exec("INSERT INTO workouts (exercise_id, date) VALUES (?,?)", workout.ExerciseID, workout.Date)
	if err != nil {
		return 0, fmt.Errorf("addWorkout %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addWorkout: %v", err)
	}
	return id, nil
}

func getWorkouts() ([]Workout, error) {
	rows, err := database.DB.Query("SELECT * FROM workouts")
	if err != nil {
		return nil, fmt.Errorf("addWorkout: %v", err)
	}
	var workouts []Workout

	for rows.Next() {
		var workout Workout
		if err := rows.Scan(&workout.ID, &workout.ExerciseID, &workout.Date); err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		workouts = append(workouts, workout)
	}
	return workouts, nil
}

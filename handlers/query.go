package handlers

import ( 
	"github.com/gopsdv/lightweight/database"
	"fmt"
)

func AddExercise(exercise database.Exercise) (int64, error) {
    result, err := database.DB.Exec("INSERT INTO exercises (name) VALUES (?)", exercise.Name )
    if err != nil {
        return 0, fmt.Errorf("addExercise: %v", err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addExercise: %v", err)
    }
    return id, nil
}


func getExercises() (any, error) {
    rows, err := database.DB.Query("SELECT * FROM exercises")
    if err != nil {
        return 0, fmt.Errorf("addExercise: %v", err)
    }
    var exercises []database.Exercise

	for rows.Next() {
        var exercise database.Exercise
        if err := rows.Scan(&exercise.ID, &exercise.Name); err != nil {
            return nil, fmt.Errorf("%v", err)
        }
        exercises = append(exercises, exercise)
    }
    return exercises, nil
}
package sets

import (
	"fmt"
	"github.com/gopsdv/lightweight/database"
)

func AddSet(set Set) (int64, error) {
	result, err := database.DB.Exec("INSERT INTO sets (workout_id, set_num, weight, reps, preps, rir) VALUES (?,?,?,?,?,?)", set.WorkoutID, set.SetNum, set.Weight, set.Reps, set.PartialReps, set.RIR)
	if err != nil {
		return 0, fmt.Errorf("addSet: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addSet: %v", err)
	}
	return id, nil
}

func getSets() ([]Set, error) {
	rows, err := database.DB.Query("SELECT * FROM sets")
	if err != nil {
		return nil, fmt.Errorf("addSet: %v", err)
	}
	var sets []Set

	for rows.Next() {
		var set Set
		if err := rows.Scan(&set.ID, &set.WorkoutID, &set.SetNum, &set.Weight, &set.Reps, &set.PartialReps, &set.RIR); err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		sets = append(sets, set)
	}
	return sets, nil
}

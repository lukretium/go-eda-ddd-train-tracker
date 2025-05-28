package domain

import (
	"time"

	"github.com/google/uuid"
)

type Exercise struct {
	Name   string
	Sets   int
	Reps   int
	Weight float64
}

type Workout struct {
	ID        string
	UserID    string
	Date      time.Time
	Exercises []Exercise
}

func NewWorkout(id, userID string, date time.Time, exercises []Exercise) (*Workout, WorkoutLoggedEvent) {
	if id == "" {
		id = uuid.New().String()
	}
	w := &Workout{
		ID:        id,
		UserID:    userID,
		Date:      date,
		Exercises: exercises,
	}
	event := WorkoutLoggedEvent{
		UserID:    userID,
		WorkoutID: id,
		Date:      date,
		Exercises: exercises,
	}
	return w, event
}

package domain

import "time"

type WorkoutLoggedEvent struct {
	UserID    string
	WorkoutID string
	Date      time.Time
	Exercises []Exercise
}

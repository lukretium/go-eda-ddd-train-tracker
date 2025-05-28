package application

import (
	"time"

	"workout/domain"
)

type LogWorkoutCommand struct {
	UserID    string
	Date      time.Time
	Exercises []domain.Exercise
}

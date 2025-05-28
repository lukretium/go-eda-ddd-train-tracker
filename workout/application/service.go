package application

import (
	"workout/domain"
)

type WorkoutService struct {
	repo      domain.WorkoutRepository
	publisher interface{} // placeholder, not used
}

func NewWorkoutService(repo domain.WorkoutRepository, publisher interface{}) *WorkoutService {
	return &WorkoutService{repo: repo, publisher: publisher}
}

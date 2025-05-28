package query

import "workout/domain"

type ListWorkoutsByUserHandler struct {
	Repo domain.WorkoutRepository
}

func NewListWorkoutsByUserHandler(repo domain.WorkoutRepository) *ListWorkoutsByUserHandler {
	return &ListWorkoutsByUserHandler{Repo: repo}
}

func (h *ListWorkoutsByUserHandler) Handle(userID string) ([]*domain.Workout, error) {
	return h.Repo.ListByUserID(userID)
}

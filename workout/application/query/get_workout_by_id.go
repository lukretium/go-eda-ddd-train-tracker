package query

import "workout/domain"

type GetWorkoutByIDHandler struct {
	Repo domain.WorkoutRepository
}

func NewGetWorkoutByIDHandler(repo domain.WorkoutRepository) *GetWorkoutByIDHandler {
	return &GetWorkoutByIDHandler{Repo: repo}
}

func (h *GetWorkoutByIDHandler) Handle(id string) (*domain.Workout, error) {
	return h.Repo.FindByID(id)
}

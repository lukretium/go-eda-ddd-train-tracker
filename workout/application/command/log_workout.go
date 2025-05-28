package command

import (
	"workout/application"
	"workout/application/port"
	"workout/domain"
)

type LogWorkoutHandler struct {
	Repo      domain.WorkoutRepository
	Publisher port.EventPublisher
}

func NewLogWorkoutHandler(repo domain.WorkoutRepository, publisher port.EventPublisher) *LogWorkoutHandler {
	return &LogWorkoutHandler{Repo: repo, Publisher: publisher}
}

func (h *LogWorkoutHandler) Handle(cmd application.LogWorkoutCommand) error {
	workout, event := domain.NewWorkout("", cmd.UserID, cmd.Date, cmd.Exercises)
	if err := h.Repo.Save(workout); err != nil {
		return err
	}
	if err := h.Publisher.Publish(event); err != nil {
		return err
	}
	return nil
}

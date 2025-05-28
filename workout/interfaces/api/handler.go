package api

import (
	"encoding/json"
	"net/http"
	"time"

	"workout/application"
	"workout/application/command"
	"workout/application/query"
	"workout/domain"
)

type LogWorkoutRequest struct {
	UserID    string            `json:"userId"`
	Date      string            `json:"date"`
	Exercises []domain.Exercise `json:"exercises"`
}

type WorkoutHandler struct {
	LogWorkoutHandler     *command.LogWorkoutHandler
	GetWorkoutByIDHandler *query.GetWorkoutByIDHandler
}

func (h *WorkoutHandler) LogWorkout(w http.ResponseWriter, r *http.Request) {
	var req LogWorkoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		http.Error(w, "invalid date format", http.StatusBadRequest)
		return
	}
	cmd := application.LogWorkoutCommand{
		UserID:    req.UserID,
		Date:      parsedDate,
		Exercises: req.Exercises,
	}
	if err := h.LogWorkoutHandler.Handle(cmd); err != nil {
		http.Error(w, "failed to log workout", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *WorkoutHandler) GetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/workouts/"):] // naive extraction, assumes /workouts/{id}
	if id == "" {
		http.Error(w, "missing workout id", http.StatusBadRequest)
		return
	}
	workout, err := h.GetWorkoutByIDHandler.Handle(id)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if workout == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workout)
}

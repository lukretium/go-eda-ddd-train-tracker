package api

import (
	"encoding/json"
	"net/http"
	"time"

	"workout/application"
	"workout/domain"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type LogWorkoutRequest struct {
	UserID    string            `json:"userId"`
	Date      string            `json:"date"`
	Exercises []domain.Exercise `json:"exercises"`
}

type LogWorkoutHandler struct {
	Handler func(w http.ResponseWriter, r *http.Request)
}

func NewLogWorkoutHandler(logWorkoutHandler func(w http.ResponseWriter, r *http.Request)) *LogWorkoutHandler {
	return &LogWorkoutHandler{Handler: logWorkoutHandler}
}

func LogWorkoutHandlerFunc(logWorkoutHandlerFunc func(cmd application.LogWorkoutCommand) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LogWorkoutRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Code: "invalid_json", Message: "invalid request body"})
			return
		}
		if verr := validateLogWorkoutRequest(req); verr != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(verr)
			return
		}
		parsedDate, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Code: "invalid_date_format", Message: "date must be YYYY-MM-DD"})
			return
		}
		cmd := application.LogWorkoutCommand{
			UserID:    req.UserID,
			Date:      parsedDate,
			Exercises: req.Exercises,
		}
		if err := logWorkoutHandlerFunc(cmd); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Code: "internal_error", Message: "failed to log workout"})
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func validateLogWorkoutRequest(req LogWorkoutRequest) *ErrorResponse {
	if req.UserID == "" {
		return &ErrorResponse{Code: "invalid_user_id", Message: "userId is required"}
	}
	if req.Date == "" {
		return &ErrorResponse{Code: "invalid_date", Message: "date is required"}
	}
	if len(req.Exercises) == 0 {
		return &ErrorResponse{Code: "invalid_exercises", Message: "at least one exercise is required"}
	}
	for i, ex := range req.Exercises {
		if ex.Name == "" {
			return &ErrorResponse{Code: "invalid_exercise_name", Message: "exercise name is required (index " + string(i) + ")"}
		}
		if ex.Sets < 0 || ex.Reps < 0 || ex.Weight < 0 {
			return &ErrorResponse{Code: "invalid_exercise_value", Message: "sets, reps, and weight must be non-negative (index " + string(i) + ")"}
		}
	}
	return nil
}

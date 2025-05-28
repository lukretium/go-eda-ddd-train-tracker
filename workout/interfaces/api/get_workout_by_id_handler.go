package api

import (
	"encoding/json"
	"net/http"

	"workout/application/query"
)

func GetWorkoutByIDHandlerFunc(getWorkoutByIDHandler *query.GetWorkoutByIDHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/workouts/"):] // naive extraction, assumes /workouts/{id}
		if id == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Code: "missing_workout_id", Message: "workout id is required"})
			return
		}
		workout, err := getWorkoutByIDHandler.Handle(id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Code: "internal_error", Message: "internal error"})
			return
		}
		if workout == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorResponse{Code: "not_found", Message: "workout not found"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(workout)
	}
}

package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"workout/application/query"
)

func ListWorkoutsByUserHandlerFunc(listHandler *query.ListWorkoutsByUserHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Expect path: /users/{userId}/workouts
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 4 || parts[1] != "users" || parts[3] != "workouts" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Code: "invalid_path", Message: "expected /users/{userId}/workouts"})
			return
		}
		userID := parts[2]
		if userID == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Code: "missing_user_id", Message: "userId is required"})
			return
		}
		workouts, err := listHandler.Handle(userID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Code: "internal_error", Message: "internal error"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(workouts)
	}
}

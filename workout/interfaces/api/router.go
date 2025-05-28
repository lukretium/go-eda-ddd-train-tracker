package api

import (
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func NewRouter(handler *WorkoutHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/workouts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.LogWorkout(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})
	mux.HandleFunc("/workouts/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetWorkoutByID(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})
	return withCORS(mux)
}

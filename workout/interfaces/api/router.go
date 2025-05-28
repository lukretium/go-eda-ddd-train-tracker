package api

import (
	"net/http"
	"workout/application/command"
	"workout/application/query"
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

func NewRouter(
	logWorkoutHandler *command.LogWorkoutHandler,
	getWorkoutByIDHandler *query.GetWorkoutByIDHandler,
	listWorkoutsByUserHandler *query.ListWorkoutsByUserHandler,
) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/workouts", LogWorkoutHandlerFunc(logWorkoutHandler.Handle))
	mux.HandleFunc("/workouts/", GetWorkoutByIDHandlerFunc(getWorkoutByIDHandler))
	mux.HandleFunc("/users/", ListWorkoutsByUserHandlerFunc(listWorkoutsByUserHandler))
	return withCORS(mux)
}

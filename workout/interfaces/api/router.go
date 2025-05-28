package api

import (
	"net/http"
	"time"
	"workout/application/command"
	"workout/application/query"

	"github.com/rs/zerolog/log"
)

type Router struct {
	mux *http.ServeMux
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := &responseWriter{ResponseWriter: w, status: 200}
		next.ServeHTTP(ww, r)
		duration := time.Since(start)
		log.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Int("status", ww.status).
			Dur("duration", duration).
			Msg("http request")
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
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
	return withCORS(loggingMiddleware(mux))
}

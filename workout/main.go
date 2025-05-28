package main

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"workout/application/command"
	"workout/application/query"
	"workout/infrastructure"
	"workout/interfaces/api"
)

func main() {
	// Use zerolog ConsoleWriter for pretty, colorized logs in the terminal
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	repo := infrastructure.NewMemoryWorkoutRepository()
	publisher := &infrastructure.StdoutPublisher{}
	logWorkoutHandler := command.NewLogWorkoutHandler(repo, publisher)
	getWorkoutByIDHandler := query.NewGetWorkoutByIDHandler(repo)
	listWorkoutsByUserHandler := query.NewListWorkoutsByUserHandler(repo)

	router := api.NewRouter(logWorkoutHandler, getWorkoutByIDHandler, listWorkoutsByUserHandler)

	log.Info().Msg("Listening on :8080")
	log.Fatal().Err(http.ListenAndServe(":8080", router)).Msg("")
}

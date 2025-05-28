package main

import (
	"log"
	"net/http"

	"workout/application/command"
	"workout/application/query"
	"workout/infrastructure"
	"workout/interfaces/api"
)

func main() {
	repo := infrastructure.NewMemoryWorkoutRepository()
	publisher := &infrastructure.StdoutPublisher{}
	logWorkoutHandler := command.NewLogWorkoutHandler(repo, publisher)
	getWorkoutByIDHandler := query.NewGetWorkoutByIDHandler(repo)

	handler := &api.WorkoutHandler{
		LogWorkoutHandler:     logWorkoutHandler,
		GetWorkoutByIDHandler: getWorkoutByIDHandler,
	}

	router := api.NewRouter(handler)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

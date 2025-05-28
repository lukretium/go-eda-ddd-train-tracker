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
	listWorkoutsByUserHandler := query.NewListWorkoutsByUserHandler(repo)

	router := api.NewRouter(logWorkoutHandler, getWorkoutByIDHandler, listWorkoutsByUserHandler)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

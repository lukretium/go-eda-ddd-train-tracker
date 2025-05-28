# Workout Service

A minimal, event-driven Workout Service for a Fitness Coaching App, following Clean Architecture and DDD principles.

## Features
- Log workouts via HTTP API
- Emits domain events (stdout for now)
- In-memory storage (easy to swap for a real DB)
- Modular, testable, and ready for extension

## Architecture
- **Domain**: Business logic, aggregates, events
- **Application**: Use cases, command handling
- **Infrastructure**: Storage, event publishing (in-memory, stdout)
- **API**: HTTP handlers and routing

## Running the Service

1. **Install Go (>=1.21)**
2. Run the service:
   ```sh
   cd workout
   go run main.go
   ```
3. The service will listen on `http://localhost:8080`

## Example Usage

### Log a Workout
```sh
curl -X POST http://localhost:8080/workouts \
  -H "Content-Type: application/json" \
  -d '{
    "userId": "abc123",
    "date": "2025-05-27",
    "exercises": [
      { "name": "Squat", "sets": 3, "reps": 10, "weight": 80 }
    ]
  }'
```

### Expected Output
- HTTP 200 OK on success
- The event will be printed to stdout, e.g.:
  ```json
  {"UserID":"abc123","WorkoutID":"","Date":"2025-05-27T00:00:00Z","Exercises":[{"Name":"Squat","Sets":3,"Reps":10,"Weight":80}]}
  ```

## Extending the Service
- Swap in a real database by implementing the `WorkoutRepository` interface
- Add more endpoints in `interfaces/api/router.go`
- Replace the event publisher with Kafka, NATS, etc.

## Project Structure
```
domain/         # Business logic, aggregates, events
application/    # Use cases, command handling
infrastructure/ # Storage, event publishing
interfaces/api/ # HTTP handlers and routing
main.go         # Service entrypoint
``` 
# Train Tracker Fitness Coaching Platform

This repository contains a modular, event-driven fitness coaching platform designed with Clean Architecture and Domain-Driven Design (DDD) principles. The system is composed of several independent services that communicate via events, enabling robust progress tracking, extensibility, and clear separation of concerns.

## 📋 Service Rules & Documentation Standards

Every service in this platform **must** follow these rules:

- **API Documentation**: Each service exposing HTTP endpoints must provide an `api.yaml` (OpenAPI 3.x) file describing all routes, request/response schemas, and error codes. This ensures discoverability and enables client code generation and testing tools (e.g., Swagger UI, Postman).
- **Event Schema Documentation**: Each service that emits or consumes events must provide an `event-schema.yaml` (or `event-schema.json`) file. This should describe all event types, their payloads, and versioning. This is essential for:
  - Allowing other teams/services to subscribe to events safely
  - Enabling event validation and evolution
  - Supporting event-driven development and integration
- **CQRS & Use Case Isolation**: Each command or query is implemented as a dedicated handler (not a god service). This keeps the codebase modular and testable.
- **CORS**: HTTP APIs must support CORS for browser-based tools and clients.
- **Singular Package Names**: All Go packages and folders use singular names for clarity and idiomatic imports.
- **README**: Each service must have a README with setup, usage, and extension instructions.

### Example: Event Schema Documentation

A minimal `event-schema.yaml` for the Workout Service might look like:

```yaml
events:
  - name: workout.logged
    description: Emitted when a workout is logged
    payload:
      userId: string
      workoutId: string
      date: string  # ISO 8601
      exercises:
        - name: string
          sets: integer
          reps: integer
          weight: number
    version: 1
  - name: workout.deleted
    description: Emitted when a workout is deleted
    payload:
      userId: string
      workoutId: string
    version: 1
```

## 🏗️ Architecture Overview

- **Microservices**: Each core domain is implemented as a separate service.
- **Event-Driven**: Services communicate by emitting and consuming domain events (e.g., via Kafka, NATS, or other brokers).
- **CQRS**: Command and Query responsibilities are separated for scalability and clarity.
- **Projections**: Read models are built asynchronously from events for fast, tailored queries.

## 🧩 Core Services

| Service             | Responsibility                                                      |
|---------------------|---------------------------------------------------------------------|
| **User Service**    | Manages user registration, profiles, and preferences                |
| **Workout Service** | Accepts workout logs, emits workout events                          |
| **Goal Tracker**    | Handles goal setting, tracks achievement, emits goal events         |
| **Projection(s)**   | Builds materialized views (dashboards, stats, leaderboards, etc.)   |
| **Notification**    | Sends feedback, reminders, and notifications                        |
| **Event Store**     | Stores all events in an append-only log                             |
| **Message Broker**  | Delivers events between services (Kafka, NATS, etc.)                |

## ⚡ Communication Between Services

- **Event Sourcing**: All changes (e.g., workout logged, goal set) are captured as immutable events.
- **Publish/Subscribe**: Services publish events to a message broker. Other services subscribe to relevant events.
- **Loose Coupling**: Services do not call each other directly; they react to events, enabling independent scaling and evolution.
- **Projections**: Read models are updated asynchronously by consuming events, ensuring eventual consistency.

### Example Event Flow
1. **User logs a workout** via the Workout Service API.
2. **Workout Service** emits a `workout.logged` event to the broker.
3. **Goal Tracker** consumes the event, checks for goal progress, and may emit a `goal.achieved` event.
4. **Projection Service** updates dashboards and stats.
5. **Notification Service** sends feedback or reminders if needed.

## 🧰 Technologies
- **Go** for service implementation
- **Kafka/NATS** for event delivery (can be swapped as needed)
- **PostgreSQL/SQLite** for projections
- **Docker Compose** for local development
- **OpenTelemetry/Prometheus** for observability

## 🚀 Getting Started
Each service contains its own README with setup and usage instructions. Start with the Workout Service for a minimal example.

--- 
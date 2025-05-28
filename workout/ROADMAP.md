# Workout Service Roadmap

This document tracks open tasks and planned enhancements for the Workout Service.

## 🏗️ Use Cases (Commands & Queries)
- [x] List workouts for a user (`GET /users/{userId}/workouts`)
- [ ] Delete a workout (`DELETE /workouts/{id}`), emit `workout.deleted` event
- [ ] Update/edit a workout (`PUT /workouts/{id}`), emit `workout.updated` event
- [ ] Filter/query workouts by date, exercise, etc.

## ✅ Validation & Error Handling
- [x] Add request validation (required fields, no negative sets/reps/weight)
- [x] Return structured error responses (error codes, details)

## 📈 Observability
- [/] Add structured logging for requests, errors, and events
  - [ ] Use a structured logger (e.g., zerolog, logrus)
  - [ ] Log all incoming HTTP requests (method, path, status, duration)
  - [ ] Log all errors with context
  - [ ] Log all emitted events (event type, payload)
- [ ] Integrate OpenTelemetry for distributed tracing
- [ ] Expose Prometheus metrics (workouts logged, errors, etc.)

## 🧪 Testing
- [ ] Unit tests for command/query handlers, repositories, event publishers
- [ ] Integration tests for API endpoints and event flows

## 💾 Persistence
- [ ] Implement a database-backed repository (PostgreSQL or SQLite)
- [ ] Add database migrations (e.g., with golang-migrate)

## 🔔 Event Delivery
- [ ] Implement a Kafka or NATS publisher (in addition to stdout)
- [ ] Ensure event consumers are idempotent (safe to replay events)

## 🛠️ API Enhancements
- [ ] Add pagination to listing endpoints
- [ ] Improve OpenAPI spec (tags, examples, authentication)

## 🔒 Security
- [ ] Require and validate JWTs or API keys for endpoints
- [ ] Input sanitization to prevent injection attacks

## 📚 Documentation
- [ ] Add more examples to README and api.yaml
- [ ] Document event schema evolution and versioning

## 🔌 Extensibility
- [ ] Add webhook support for external event subscribers
- [ ] Add gRPC or GraphQL API for advanced/internal integrations

---
**Prioritize tasks based on product needs, team capacity, and feedback!** 
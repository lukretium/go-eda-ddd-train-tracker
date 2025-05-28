package port

type EventPublisher interface {
	Publish(event interface{}) error
}

package infrastructure

import (
	"encoding/json"
	"fmt"
	"workout/application/port"
)

type StdoutPublisher struct{}

func (p *StdoutPublisher) Publish(event interface{}) error {
	b, err := json.Marshal(event)
	if err != nil {
		return err
	}
	fmt.Println("EVENT:", string(b))
	return nil
}

var _ port.EventPublisher = (*StdoutPublisher)(nil)

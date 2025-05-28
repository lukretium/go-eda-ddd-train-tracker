package infrastructure

import (
	"encoding/json"
	"workout/application/port"

	"github.com/rs/zerolog/log"
)

type StdoutPublisher struct{}

func (p *StdoutPublisher) Publish(event interface{}) error {
	b, err := json.Marshal(event)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal event for logging")
		return err
	}
	log.Info().RawJSON("event", b).Msg("event emitted")
	return nil
}

var _ port.EventPublisher = (*StdoutPublisher)(nil)

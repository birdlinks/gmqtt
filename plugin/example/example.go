package example

import (
	"github.com/birdlinks/gmqtt/internal/events"
	"github.com/birdlinks/gmqtt/internal/log"
	"github.com/birdlinks/gmqtt/internal/server"
)

func New() (server.Plugin, error) {
	return &Example{}, nil
}

type Example struct {
}

func (h *Example) Name() string {
	return "example-plugin"
}

func (h *Example) Version() string {
	return "0.0.1"
}

func (h *Example) Init(s server.Server) error {
	log.Info("example plugin started...")
	return nil
}

func (h *Example) Close() error {
	log.Info("example plugin stopped...")
	return nil
}

func (h *Example) Hook() server.HookWrapper {
	return server.HookWrapper{
		OnMessage: func(message events.OnMessage) events.OnMessage {
			return func(client events.Client, packet events.Packet) (events.Packet, error) {
				log.Info("message received",
					log.Any("msg", string(packet.Payload)),
				)
				return packet, nil
			}
		},
	}
}

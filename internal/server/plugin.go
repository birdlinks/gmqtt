package server

import (
	"github.com/birdlinks/gmqtt/internal/events"
)

type HookWrapper struct {
	OnProcessMessage events.OnProcessMessageWrapper
	OnMessage        events.OnMessageWrapper
	OnError          events.OnErrorWrapper
	OnConnect        events.OnConnectWrapper
	OnDisconnect     events.OnDisconnectWrapper
	OnSubscribe      events.OnSubscribeWrapper
	OnUnsubscribe    events.OnUnsubscribeWrapper
}

type NewPlugin func() (Plugin, error)

type Plugin interface {
	// Name returns the name of the plugin.
	Name() string
	// Version returns the version of the plugin.
	Version() string
	// Init initializes the plugin.
	Init(server Server) error
	// Close closes the plugin.
	Close() error
	// Hook returns the hook of the plugin.
	Hook() HookWrapper
}

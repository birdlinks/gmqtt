package plugin

import "github.com/birdlinks/gmqtt/internal/events"

type HookWrapper struct {
	OnProcessMessage events.OnProcessMessage
	OnMessage        events.OnMessage
	OnError          events.OnError
	OnConnect        events.OnConnect
	OnDisconnect     events.OnDisconnect
	OnSubscribe      events.OnSubscribe
	OnUnsubscribe    events.OnUnsubscribe
}

type NewPlugin func() (Plugin, error)

type Plugin interface {
	// Name returns the name of the plugin.
	Name() string
	// Version returns the version of the plugin.
	Version() string
	// Init initializes the plugin.
	Init() error
	// Close closes the plugin.
	Close() error
	// Hook returns the hook of the plugin.
	Hook() HookWrapper
}

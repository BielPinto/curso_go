package events

import "time"

type EventInterface interface {
	GetName() string
	GetDataTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Habdle(event EventInterface)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventHandlerInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}

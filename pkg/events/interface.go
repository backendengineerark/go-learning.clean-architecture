package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	getDateTime() time.Time
	GetPayload() interface{}
	SetPayload(interface{})
}

type EventHandlerInterface interface {
	Handle(event EventHandlerInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) error
	Clear()
}

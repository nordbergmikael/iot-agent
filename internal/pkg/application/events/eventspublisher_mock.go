// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package events

import (
	"context"
	"github.com/diwise/iot-agent/internal/pkg/application/conversion"
	"sync"
)

// Ensure, that EventPublisherMock does implement EventPublisher.
// If this is not the case, regenerate this file with moq.
var _ EventPublisher = &EventPublisherMock{}

// EventPublisherMock is a mock implementation of EventPublisher.
//
// 	func TestSomethingThatUsesEventPublisher(t *testing.T) {
//
// 		// make and configure a mocked EventPublisher
// 		mockedEventPublisher := &EventPublisherMock{
// 			PublishFunc: func(ctx context.Context, msg conversion.InternalMessage) error {
// 				panic("mock out the Publish method")
// 			},
// 			StartFunc: func() error {
// 				panic("mock out the Start method")
// 			},
// 			StopFunc: func() error {
// 				panic("mock out the Stop method")
// 			},
// 		}
//
// 		// use mockedEventPublisher in code that requires EventPublisher
// 		// and then make assertions.
//
// 	}
type EventPublisherMock struct {
	// PublishFunc mocks the Publish method.
	PublishFunc func(ctx context.Context, msg conversion.InternalMessage) error

	// StartFunc mocks the Start method.
	StartFunc func() error

	// StopFunc mocks the Stop method.
	StopFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// Publish holds details about calls to the Publish method.
		Publish []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Msg is the msg argument value.
			Msg conversion.InternalMessage
		}
		// Start holds details about calls to the Start method.
		Start []struct {
		}
		// Stop holds details about calls to the Stop method.
		Stop []struct {
		}
	}
	lockPublish sync.RWMutex
	lockStart   sync.RWMutex
	lockStop    sync.RWMutex
}

// Publish calls PublishFunc.
func (mock *EventPublisherMock) Publish(ctx context.Context, msg conversion.InternalMessage) error {
	if mock.PublishFunc == nil {
		panic("EventPublisherMock.PublishFunc: method is nil but EventPublisher.Publish was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Msg conversion.InternalMessage
	}{
		Ctx: ctx,
		Msg: msg,
	}
	mock.lockPublish.Lock()
	mock.calls.Publish = append(mock.calls.Publish, callInfo)
	mock.lockPublish.Unlock()
	return mock.PublishFunc(ctx, msg)
}

// PublishCalls gets all the calls that were made to Publish.
// Check the length with:
//     len(mockedEventPublisher.PublishCalls())
func (mock *EventPublisherMock) PublishCalls() []struct {
	Ctx context.Context
	Msg conversion.InternalMessage
} {
	var calls []struct {
		Ctx context.Context
		Msg conversion.InternalMessage
	}
	mock.lockPublish.RLock()
	calls = mock.calls.Publish
	mock.lockPublish.RUnlock()
	return calls
}

// Start calls StartFunc.
func (mock *EventPublisherMock) Start() error {
	if mock.StartFunc == nil {
		panic("EventPublisherMock.StartFunc: method is nil but EventPublisher.Start was just called")
	}
	callInfo := struct {
	}{}
	mock.lockStart.Lock()
	mock.calls.Start = append(mock.calls.Start, callInfo)
	mock.lockStart.Unlock()
	return mock.StartFunc()
}

// StartCalls gets all the calls that were made to Start.
// Check the length with:
//     len(mockedEventPublisher.StartCalls())
func (mock *EventPublisherMock) StartCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStart.RLock()
	calls = mock.calls.Start
	mock.lockStart.RUnlock()
	return calls
}

// Stop calls StopFunc.
func (mock *EventPublisherMock) Stop() error {
	if mock.StopFunc == nil {
		panic("EventPublisherMock.StopFunc: method is nil but EventPublisher.Stop was just called")
	}
	callInfo := struct {
	}{}
	mock.lockStop.Lock()
	mock.calls.Stop = append(mock.calls.Stop, callInfo)
	mock.lockStop.Unlock()
	return mock.StopFunc()
}

// StopCalls gets all the calls that were made to Stop.
// Check the length with:
//     len(mockedEventPublisher.StopCalls())
func (mock *EventPublisherMock) StopCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStop.RLock()
	calls = mock.calls.Stop
	mock.lockStop.RUnlock()
	return calls
}

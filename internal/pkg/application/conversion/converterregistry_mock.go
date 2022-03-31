// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package conversion

import (
	"context"
	"sync"
)

// Ensure, that ConverterRegistryMock does implement ConverterRegistry.
// If this is not the case, regenerate this file with moq.
var _ ConverterRegistry = &ConverterRegistryMock{}

// ConverterRegistryMock is a mock implementation of ConverterRegistry.
//
// 	func TestSomethingThatUsesConverterRegistry(t *testing.T) {
//
// 		// make and configure a mocked ConverterRegistry
// 		mockedConverterRegistry := &ConverterRegistryMock{
// 			DesignateConvertersFunc: func(ctx context.Context, types []string) []MessageConverterFunc {
// 				panic("mock out the DesignateConverters method")
// 			},
// 		}
//
// 		// use mockedConverterRegistry in code that requires ConverterRegistry
// 		// and then make assertions.
//
// 	}
type ConverterRegistryMock struct {
	// DesignateConvertersFunc mocks the DesignateConverters method.
	DesignateConvertersFunc func(ctx context.Context, types []string) []MessageConverterFunc

	// calls tracks calls to the methods.
	calls struct {
		// DesignateConverters holds details about calls to the DesignateConverters method.
		DesignateConverters []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Types is the types argument value.
			Types []string
		}
	}
	lockDesignateConverters sync.RWMutex
}

// DesignateConverters calls DesignateConvertersFunc.
func (mock *ConverterRegistryMock) DesignateConverters(ctx context.Context, types []string) []MessageConverterFunc {
	if mock.DesignateConvertersFunc == nil {
		panic("ConverterRegistryMock.DesignateConvertersFunc: method is nil but ConverterRegistry.DesignateConverters was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Types []string
	}{
		Ctx:   ctx,
		Types: types,
	}
	mock.lockDesignateConverters.Lock()
	mock.calls.DesignateConverters = append(mock.calls.DesignateConverters, callInfo)
	mock.lockDesignateConverters.Unlock()
	return mock.DesignateConvertersFunc(ctx, types)
}

// DesignateConvertersCalls gets all the calls that were made to DesignateConverters.
// Check the length with:
//     len(mockedConverterRegistry.DesignateConvertersCalls())
func (mock *ConverterRegistryMock) DesignateConvertersCalls() []struct {
	Ctx   context.Context
	Types []string
} {
	var calls []struct {
		Ctx   context.Context
		Types []string
	}
	mock.lockDesignateConverters.RLock()
	calls = mock.calls.DesignateConverters
	mock.lockDesignateConverters.RUnlock()
	return calls
}
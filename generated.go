// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package main

import (
	"github.com/johanbrandhorst/moq-vendor-bug/vendor/github.com/lib/pq"
	"sync"
)

var (
	lockMyInterfaceMockDoSomething sync.RWMutex
)

// MyInterfaceMock is a mock implementation of MyInterface.
//
//     func TestSomethingThatUsesMyInterface(t *testing.T) {
//
//         // make and configure a mocked MyInterface
//         mockedMyInterface := &MyInterfaceMock{
//             DoSomethingFunc: func(in1 pq.ByteaArray) pq.BoolArray {
// 	               panic("TODO: mock out the DoSomething method")
//             },
//         }
//
//         // TODO: use mockedMyInterface in code that requires MyInterface
//         //       and then make assertions.
//
//     }
type MyInterfaceMock struct {
	// DoSomethingFunc mocks the DoSomething method.
	DoSomethingFunc func(in1 pq.ByteaArray) pq.BoolArray

	// calls tracks calls to the methods.
	calls struct {
		// DoSomething holds details about calls to the DoSomething method.
		DoSomething []struct {
			// In1 is the in1 argument value.
			In1 pq.ByteaArray
		}
	}
}

// DoSomething calls DoSomethingFunc.
func (mock *MyInterfaceMock) DoSomething(in1 pq.ByteaArray) pq.BoolArray {
	if mock.DoSomethingFunc == nil {
		panic("moq: MyInterfaceMock.DoSomethingFunc is nil but MyInterface.DoSomething was just called")
	}
	callInfo := struct {
		In1 pq.ByteaArray
	}{
		In1: in1,
	}
	lockMyInterfaceMockDoSomething.Lock()
	mock.calls.DoSomething = append(mock.calls.DoSomething, callInfo)
	lockMyInterfaceMockDoSomething.Unlock()
	return mock.DoSomethingFunc(in1)
}

// DoSomethingCalls gets all the calls that were made to DoSomething.
// Check the length with:
//     len(mockedMyInterface.DoSomethingCalls())
func (mock *MyInterfaceMock) DoSomethingCalls() []struct {
	In1 pq.ByteaArray
} {
	var calls []struct {
		In1 pq.ByteaArray
	}
	lockMyInterfaceMockDoSomething.RLock()
	calls = mock.calls.DoSomething
	lockMyInterfaceMockDoSomething.RUnlock()
	return calls
}
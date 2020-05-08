// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mongoke

import (
	"sync"
)

var (
	lockDatabaseInterfaceMockFindMany sync.RWMutex
	lockDatabaseInterfaceMockFindOne  sync.RWMutex
)

// Ensure, that DatabaseInterfaceMock does implement DatabaseInterface.
// If this is not the case, regenerate this file with moq.
var _ DatabaseInterface = &DatabaseInterfaceMock{}

// DatabaseInterfaceMock is a mock implementation of DatabaseInterface.
//
//     func TestSomethingThatUsesDatabaseInterface(t *testing.T) {
//
//         // make and configure a mocked DatabaseInterface
//         mockedDatabaseInterface := &DatabaseInterfaceMock{
//             FindManyFunc: func(p FindManyParams) (Connection, error) {
// 	               panic("mock out the FindMany method")
//             },
//             FindOneFunc: func(p FindOneParams) (interface{}, error) {
// 	               panic("mock out the FindOne method")
//             },
//         }
//
//         // use mockedDatabaseInterface in code that requires DatabaseInterface
//         // and then make assertions.
//
//     }
type DatabaseInterfaceMock struct {
	// FindManyFunc mocks the FindMany method.
	FindManyFunc func(p FindManyParams) (Connection, error)

	// FindOneFunc mocks the FindOne method.
	FindOneFunc func(p FindOneParams) (interface{}, error)

	// calls tracks calls to the methods.
	calls struct {
		// FindMany holds details about calls to the FindMany method.
		FindMany []struct {
			// P is the p argument value.
			P FindManyParams
		}
		// FindOne holds details about calls to the FindOne method.
		FindOne []struct {
			// P is the p argument value.
			P FindOneParams
		}
	}
}

// FindMany calls FindManyFunc.
func (mock *DatabaseInterfaceMock) FindMany(p FindManyParams) (Connection, error) {
	if mock.FindManyFunc == nil {
		panic("DatabaseInterfaceMock.FindManyFunc: method is nil but DatabaseInterface.FindMany was just called")
	}
	callInfo := struct {
		P FindManyParams
	}{
		P: p,
	}
	lockDatabaseInterfaceMockFindMany.Lock()
	mock.calls.FindMany = append(mock.calls.FindMany, callInfo)
	lockDatabaseInterfaceMockFindMany.Unlock()
	return mock.FindManyFunc(p)
}

// FindManyCalls gets all the calls that were made to FindMany.
// Check the length with:
//     len(mockedDatabaseInterface.FindManyCalls())
func (mock *DatabaseInterfaceMock) FindManyCalls() []struct {
	P FindManyParams
} {
	var calls []struct {
		P FindManyParams
	}
	lockDatabaseInterfaceMockFindMany.RLock()
	calls = mock.calls.FindMany
	lockDatabaseInterfaceMockFindMany.RUnlock()
	return calls
}

// FindOne calls FindOneFunc.
func (mock *DatabaseInterfaceMock) FindOne(p FindOneParams) (interface{}, error) {
	if mock.FindOneFunc == nil {
		panic("DatabaseInterfaceMock.FindOneFunc: method is nil but DatabaseInterface.FindOne was just called")
	}
	callInfo := struct {
		P FindOneParams
	}{
		P: p,
	}
	lockDatabaseInterfaceMockFindOne.Lock()
	mock.calls.FindOne = append(mock.calls.FindOne, callInfo)
	lockDatabaseInterfaceMockFindOne.Unlock()
	return mock.FindOneFunc(p)
}

// FindOneCalls gets all the calls that were made to FindOne.
// Check the length with:
//     len(mockedDatabaseInterface.FindOneCalls())
func (mock *DatabaseInterfaceMock) FindOneCalls() []struct {
	P FindOneParams
} {
	var calls []struct {
		P FindOneParams
	}
	lockDatabaseInterfaceMockFindOne.RLock()
	calls = mock.calls.FindOne
	lockDatabaseInterfaceMockFindOne.RUnlock()
	return calls
}

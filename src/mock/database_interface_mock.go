// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/remorses/mongoke/src"
	"sync"
)

var (
	lockDatabaseInterfaceMockFindMany   sync.RWMutex
	lockDatabaseInterfaceMockInsertMany sync.RWMutex
	lockDatabaseInterfaceMockUpdateOne  sync.RWMutex
)

// Ensure, that DatabaseInterfaceMock does implement mongoke.DatabaseInterface.
// If this is not the case, regenerate this file with moq.
var _ mongoke.DatabaseInterface = &DatabaseInterfaceMock{}

// DatabaseInterfaceMock is a mock implementation of mongoke.DatabaseInterface.
//
//     func TestSomethingThatUsesDatabaseInterface(t *testing.T) {
//
//         // make and configure a mocked mongoke.DatabaseInterface
//         mockedDatabaseInterface := &DatabaseInterfaceMock{
//             FindManyFunc: func(ctx context.Context, p mongoke.FindManyParams) ([]mongoke.Map, error) {
// 	               panic("mock out the FindMany method")
//             },
//             InsertManyFunc: func(ctx context.Context, p mongoke.InsertManyParams) ([]mongoke.Map, error) {
// 	               panic("mock out the InsertMany method")
//             },
//             UpdateOneFunc: func(ctx context.Context, p mongoke.UpdateOneParams) (mongoke.NodeMutationPayload, error) {
// 	               panic("mock out the UpdateOne method")
//             },
//         }
//
//         // use mockedDatabaseInterface in code that requires mongoke.DatabaseInterface
//         // and then make assertions.
//
//     }
type DatabaseInterfaceMock struct {
	// FindManyFunc mocks the FindMany method.
	FindManyFunc func(ctx context.Context, p mongoke.FindManyParams) ([]mongoke.Map, error)

	// InsertManyFunc mocks the InsertMany method.
	InsertManyFunc func(ctx context.Context, p mongoke.InsertManyParams) ([]mongoke.Map, error)

	// UpdateOneFunc mocks the UpdateOne method.
	UpdateOneFunc func(ctx context.Context, p mongoke.UpdateOneParams) (mongoke.NodeMutationPayload, error)

	// calls tracks calls to the methods.
	calls struct {
		// FindMany holds details about calls to the FindMany method.
		FindMany []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// P is the p argument value.
			P mongoke.FindManyParams
		}
		// InsertMany holds details about calls to the InsertMany method.
		InsertMany []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// P is the p argument value.
			P mongoke.InsertManyParams
		}
		// UpdateOne holds details about calls to the UpdateOne method.
		UpdateOne []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// P is the p argument value.
			P mongoke.UpdateOneParams
		}
	}
}

// FindMany calls FindManyFunc.
func (mock *DatabaseInterfaceMock) FindMany(ctx context.Context, p mongoke.FindManyParams) ([]mongoke.Map, error) {
	if mock.FindManyFunc == nil {
		panic("DatabaseInterfaceMock.FindManyFunc: method is nil but DatabaseInterface.FindMany was just called")
	}
	callInfo := struct {
		Ctx context.Context
		P   mongoke.FindManyParams
	}{
		Ctx: ctx,
		P:   p,
	}
	lockDatabaseInterfaceMockFindMany.Lock()
	mock.calls.FindMany = append(mock.calls.FindMany, callInfo)
	lockDatabaseInterfaceMockFindMany.Unlock()
	return mock.FindManyFunc(ctx, p)
}

// FindManyCalls gets all the calls that were made to FindMany.
// Check the length with:
//     len(mockedDatabaseInterface.FindManyCalls())
func (mock *DatabaseInterfaceMock) FindManyCalls() []struct {
	Ctx context.Context
	P   mongoke.FindManyParams
} {
	var calls []struct {
		Ctx context.Context
		P   mongoke.FindManyParams
	}
	lockDatabaseInterfaceMockFindMany.RLock()
	calls = mock.calls.FindMany
	lockDatabaseInterfaceMockFindMany.RUnlock()
	return calls
}

// InsertMany calls InsertManyFunc.
func (mock *DatabaseInterfaceMock) InsertMany(ctx context.Context, p mongoke.InsertManyParams) ([]mongoke.Map, error) {
	if mock.InsertManyFunc == nil {
		panic("DatabaseInterfaceMock.InsertManyFunc: method is nil but DatabaseInterface.InsertMany was just called")
	}
	callInfo := struct {
		Ctx context.Context
		P   mongoke.InsertManyParams
	}{
		Ctx: ctx,
		P:   p,
	}
	lockDatabaseInterfaceMockInsertMany.Lock()
	mock.calls.InsertMany = append(mock.calls.InsertMany, callInfo)
	lockDatabaseInterfaceMockInsertMany.Unlock()
	return mock.InsertManyFunc(ctx, p)
}

// InsertManyCalls gets all the calls that were made to InsertMany.
// Check the length with:
//     len(mockedDatabaseInterface.InsertManyCalls())
func (mock *DatabaseInterfaceMock) InsertManyCalls() []struct {
	Ctx context.Context
	P   mongoke.InsertManyParams
} {
	var calls []struct {
		Ctx context.Context
		P   mongoke.InsertManyParams
	}
	lockDatabaseInterfaceMockInsertMany.RLock()
	calls = mock.calls.InsertMany
	lockDatabaseInterfaceMockInsertMany.RUnlock()
	return calls
}

// UpdateOne calls UpdateOneFunc.
func (mock *DatabaseInterfaceMock) UpdateOne(ctx context.Context, p mongoke.UpdateOneParams) (mongoke.NodeMutationPayload, error) {
	if mock.UpdateOneFunc == nil {
		panic("DatabaseInterfaceMock.UpdateOneFunc: method is nil but DatabaseInterface.UpdateOne was just called")
	}
	callInfo := struct {
		Ctx context.Context
		P   mongoke.UpdateOneParams
	}{
		Ctx: ctx,
		P:   p,
	}
	lockDatabaseInterfaceMockUpdateOne.Lock()
	mock.calls.UpdateOne = append(mock.calls.UpdateOne, callInfo)
	lockDatabaseInterfaceMockUpdateOne.Unlock()
	return mock.UpdateOneFunc(ctx, p)
}

// UpdateOneCalls gets all the calls that were made to UpdateOne.
// Check the length with:
//     len(mockedDatabaseInterface.UpdateOneCalls())
func (mock *DatabaseInterfaceMock) UpdateOneCalls() []struct {
	Ctx context.Context
	P   mongoke.UpdateOneParams
} {
	var calls []struct {
		Ctx context.Context
		P   mongoke.UpdateOneParams
	}
	lockDatabaseInterfaceMockUpdateOne.RLock()
	calls = mock.calls.UpdateOne
	lockDatabaseInterfaceMockUpdateOne.RUnlock()
	return calls
}

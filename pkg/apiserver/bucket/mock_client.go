// Code generated by mockery v2.10.0. DO NOT EDIT.

package bucket

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

// Open provides a mock function with given fields: ctx
func (_m *MockClient) Open(ctx context.Context) (Object, error) {
	ret := _m.Called(ctx)

	var r0 Object
	if rf, ok := ret.Get(0).(func(context.Context) Object); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Object)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

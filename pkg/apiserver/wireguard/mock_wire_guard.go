// Code generated by mockery v2.10.0. DO NOT EDIT.

package wireguard

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockWireGuard is an autogenerated mock type for the WireGuard type
type MockWireGuard struct {
	mock.Mock
}

// Sync provides a mock function with given fields: ctx
func (_m *MockWireGuard) Sync(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

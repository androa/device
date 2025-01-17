// Code generated by mockery v2.10.0. DO NOT EDIT.

package bucket

import (
	io "io"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockObject is an autogenerated mock type for the Object type
type MockObject struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockObject) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LastUpdated provides a mock function with given fields:
func (_m *MockObject) LastUpdated() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// Reader provides a mock function with given fields:
func (_m *MockObject) Reader() io.Reader {
	ret := _m.Called()

	var r0 io.Reader
	if rf, ok := ret.Get(0).(func() io.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	return r0
}

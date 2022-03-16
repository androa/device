// Code generated by mockery v2.10.0. DO NOT EDIT.

package pb

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockAPIServerClient is an autogenerated mock type for the APIServerClient type
type MockAPIServerClient struct {
	mock.Mock
}

// EnrollGateway provides a mock function with given fields: ctx, in, opts
func (_m *MockAPIServerClient) EnrollGateway(ctx context.Context, in *ModifyGatewayRequest, opts ...grpc.CallOption) (*ModifyGatewayResponse, error) {
	_va := make([]any, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []any
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ModifyGatewayResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ModifyGatewayRequest, ...grpc.CallOption) *ModifyGatewayResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ModifyGatewayResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ModifyGatewayRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceConfiguration provides a mock function with given fields: ctx, in, opts
func (_m *MockAPIServerClient) GetDeviceConfiguration(ctx context.Context, in *GetDeviceConfigurationRequest, opts ...grpc.CallOption) (APIServer_GetDeviceConfigurationClient, error) {
	_va := make([]any, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []any
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 APIServer_GetDeviceConfigurationClient
	if rf, ok := ret.Get(0).(func(context.Context, *GetDeviceConfigurationRequest, ...grpc.CallOption) APIServer_GetDeviceConfigurationClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(APIServer_GetDeviceConfigurationClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetDeviceConfigurationRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGateway provides a mock function with given fields: ctx, in, opts
func (_m *MockAPIServerClient) GetGateway(ctx context.Context, in *ModifyGatewayRequest, opts ...grpc.CallOption) (*Gateway, error) {
	_va := make([]any, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []any
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *Gateway
	if rf, ok := ret.Get(0).(func(context.Context, *ModifyGatewayRequest, ...grpc.CallOption) *Gateway); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Gateway)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ModifyGatewayRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGatewayConfiguration provides a mock function with given fields: ctx, in, opts
func (_m *MockAPIServerClient) GetGatewayConfiguration(ctx context.Context, in *GetGatewayConfigurationRequest, opts ...grpc.CallOption) (APIServer_GetGatewayConfigurationClient, error) {
	_va := make([]any, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []any
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 APIServer_GetGatewayConfigurationClient
	if rf, ok := ret.Get(0).(func(context.Context, *GetGatewayConfigurationRequest, ...grpc.CallOption) APIServer_GetGatewayConfigurationClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(APIServer_GetGatewayConfigurationClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetGatewayConfigurationRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListGateways provides a mock function with given fields: ctx, in, opts
func (_m *MockAPIServerClient) ListGateways(ctx context.Context, in *ListGatewayRequest, opts ...grpc.CallOption) (APIServer_ListGatewaysClient, error) {
	_va := make([]any, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []any
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 APIServer_ListGatewaysClient
	if rf, ok := ret.Get(0).(func(context.Context, *ListGatewayRequest, ...grpc.CallOption) APIServer_ListGatewaysClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(APIServer_ListGatewaysClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ListGatewayRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, in, opts
func (_m *MockAPIServerClient) Login(ctx context.Context, in *APIServerLoginRequest, opts ...grpc.CallOption) (*APIServerLoginResponse, error) {
	_va := make([]any, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []any
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *APIServerLoginResponse
	if rf, ok := ret.Get(0).(func(context.Context, *APIServerLoginRequest, ...grpc.CallOption) *APIServerLoginResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*APIServerLoginResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *APIServerLoginRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateGateway provides a mock function with given fields: ctx, in, opts
func (_m *MockAPIServerClient) UpdateGateway(ctx context.Context, in *ModifyGatewayRequest, opts ...grpc.CallOption) (*ModifyGatewayResponse, error) {
	_va := make([]any, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []any
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ModifyGatewayResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ModifyGatewayRequest, ...grpc.CallOption) *ModifyGatewayResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ModifyGatewayResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ModifyGatewayRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// FlawLeaderElector is an autogenerated mock type for the leaderElector type
type FlawLeaderElector struct {
	mock.Mock
}

type FlawLeaderElector_Expecter struct {
	mock *mock.Mock
}

func (_m *FlawLeaderElector) EXPECT() *FlawLeaderElector_Expecter {
	return &FlawLeaderElector_Expecter{mock: &_m.Mock}
}

// IfLeader provides a mock function with given fields: ctx, fn
func (_m *FlawLeaderElector) IfLeader(ctx context.Context, fn func() error) {
	_m.Called(ctx, fn)
}

// FlawLeaderElector_IfLeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IfLeader'
type FlawLeaderElector_IfLeader_Call struct {
	*mock.Call
}

// IfLeader is a helper method to define mock.On call
//   - ctx context.Context
//   - fn func() error
func (_e *FlawLeaderElector_Expecter) IfLeader(ctx interface{}, fn interface{}) *FlawLeaderElector_IfLeader_Call {
	return &FlawLeaderElector_IfLeader_Call{Call: _e.mock.On("IfLeader", ctx, fn)}
}

func (_c *FlawLeaderElector_IfLeader_Call) Run(run func(ctx context.Context, fn func() error)) *FlawLeaderElector_IfLeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func() error))
	})
	return _c
}

func (_c *FlawLeaderElector_IfLeader_Call) Return() *FlawLeaderElector_IfLeader_Call {
	_c.Call.Return()
	return _c
}

func (_c *FlawLeaderElector_IfLeader_Call) RunAndReturn(run func(context.Context, func() error)) *FlawLeaderElector_IfLeader_Call {
	_c.Call.Return(run)
	return _c
}

// NewFlawLeaderElector creates a new instance of FlawLeaderElector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFlawLeaderElector(t interface {
	mock.TestingT
	Cleanup(func())
}) *FlawLeaderElector {
	mock := &FlawLeaderElector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

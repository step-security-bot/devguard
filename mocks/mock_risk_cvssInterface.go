// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// RiskCvssInterface is an autogenerated mock type for the cvssInterface type
type RiskCvssInterface struct {
	mock.Mock
}

type RiskCvssInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *RiskCvssInterface) EXPECT() *RiskCvssInterface_Expecter {
	return &RiskCvssInterface_Expecter{mock: &_m.Mock}
}

// EnvironmentalScore provides a mock function with given fields:
func (_m *RiskCvssInterface) EnvironmentalScore() float64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EnvironmentalScore")
	}

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// RiskCvssInterface_EnvironmentalScore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnvironmentalScore'
type RiskCvssInterface_EnvironmentalScore_Call struct {
	*mock.Call
}

// EnvironmentalScore is a helper method to define mock.On call
func (_e *RiskCvssInterface_Expecter) EnvironmentalScore() *RiskCvssInterface_EnvironmentalScore_Call {
	return &RiskCvssInterface_EnvironmentalScore_Call{Call: _e.mock.On("EnvironmentalScore")}
}

func (_c *RiskCvssInterface_EnvironmentalScore_Call) Run(run func()) *RiskCvssInterface_EnvironmentalScore_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RiskCvssInterface_EnvironmentalScore_Call) Return(_a0 float64) *RiskCvssInterface_EnvironmentalScore_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RiskCvssInterface_EnvironmentalScore_Call) RunAndReturn(run func() float64) *RiskCvssInterface_EnvironmentalScore_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: key
func (_m *RiskCvssInterface) Get(key string) (string, error) {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RiskCvssInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type RiskCvssInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - key string
func (_e *RiskCvssInterface_Expecter) Get(key interface{}) *RiskCvssInterface_Get_Call {
	return &RiskCvssInterface_Get_Call{Call: _e.mock.On("Get", key)}
}

func (_c *RiskCvssInterface_Get_Call) Run(run func(key string)) *RiskCvssInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RiskCvssInterface_Get_Call) Return(_a0 string, _a1 error) *RiskCvssInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RiskCvssInterface_Get_Call) RunAndReturn(run func(string) (string, error)) *RiskCvssInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: key, value
func (_m *RiskCvssInterface) Set(key string, value string) error {
	ret := _m.Called(key, value)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RiskCvssInterface_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type RiskCvssInterface_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - key string
//   - value string
func (_e *RiskCvssInterface_Expecter) Set(key interface{}, value interface{}) *RiskCvssInterface_Set_Call {
	return &RiskCvssInterface_Set_Call{Call: _e.mock.On("Set", key, value)}
}

func (_c *RiskCvssInterface_Set_Call) Run(run func(key string, value string)) *RiskCvssInterface_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *RiskCvssInterface_Set_Call) Return(_a0 error) *RiskCvssInterface_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RiskCvssInterface_Set_Call) RunAndReturn(run func(string, string) error) *RiskCvssInterface_Set_Call {
	_c.Call.Return(run)
	return _c
}

// TemporalScore provides a mock function with given fields:
func (_m *RiskCvssInterface) TemporalScore() float64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TemporalScore")
	}

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// RiskCvssInterface_TemporalScore_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TemporalScore'
type RiskCvssInterface_TemporalScore_Call struct {
	*mock.Call
}

// TemporalScore is a helper method to define mock.On call
func (_e *RiskCvssInterface_Expecter) TemporalScore() *RiskCvssInterface_TemporalScore_Call {
	return &RiskCvssInterface_TemporalScore_Call{Call: _e.mock.On("TemporalScore")}
}

func (_c *RiskCvssInterface_TemporalScore_Call) Run(run func()) *RiskCvssInterface_TemporalScore_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RiskCvssInterface_TemporalScore_Call) Return(_a0 float64) *RiskCvssInterface_TemporalScore_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RiskCvssInterface_TemporalScore_Call) RunAndReturn(run func() float64) *RiskCvssInterface_TemporalScore_Call {
	_c.Call.Return(run)
	return _c
}

// Vector provides a mock function with given fields:
func (_m *RiskCvssInterface) Vector() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Vector")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RiskCvssInterface_Vector_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Vector'
type RiskCvssInterface_Vector_Call struct {
	*mock.Call
}

// Vector is a helper method to define mock.On call
func (_e *RiskCvssInterface_Expecter) Vector() *RiskCvssInterface_Vector_Call {
	return &RiskCvssInterface_Vector_Call{Call: _e.mock.On("Vector")}
}

func (_c *RiskCvssInterface_Vector_Call) Run(run func()) *RiskCvssInterface_Vector_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RiskCvssInterface_Vector_Call) Return(_a0 string) *RiskCvssInterface_Vector_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RiskCvssInterface_Vector_Call) RunAndReturn(run func() string) *RiskCvssInterface_Vector_Call {
	_c.Call.Return(run)
	return _c
}

// NewRiskCvssInterface creates a new instance of RiskCvssInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRiskCvssInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *RiskCvssInterface {
	mock := &RiskCvssInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

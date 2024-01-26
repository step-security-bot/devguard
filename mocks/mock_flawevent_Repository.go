// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	flaw "github.com/l3montree-dev/flawfix/internal/core/flaw"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// FlaweventRepository is an autogenerated mock type for the Repository type
type FlaweventRepository struct {
	mock.Mock
}

type FlaweventRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *FlaweventRepository) EXPECT() *FlaweventRepository_Expecter {
	return &FlaweventRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: tx, t
func (_m *FlaweventRepository) Create(tx *gorm.DB, t *flaw.Model) error {
	ret := _m.Called(tx, t)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *flaw.Model) error); ok {
		r0 = rf(tx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlaweventRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type FlaweventRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - tx *gorm.DB
//   - t *flaw.Model
func (_e *FlaweventRepository_Expecter) Create(tx interface{}, t interface{}) *FlaweventRepository_Create_Call {
	return &FlaweventRepository_Create_Call{Call: _e.mock.On("Create", tx, t)}
}

func (_c *FlaweventRepository_Create_Call) Run(run func(tx *gorm.DB, t *flaw.Model)) *FlaweventRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*flaw.Model))
	})
	return _c
}

func (_c *FlaweventRepository_Create_Call) Return(_a0 error) *FlaweventRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FlaweventRepository_Create_Call) RunAndReturn(run func(*gorm.DB, *flaw.Model) error) *FlaweventRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: tx, id
func (_m *FlaweventRepository) Delete(tx *gorm.DB, id uuid.UUID) error {
	ret := _m.Called(tx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, uuid.UUID) error); ok {
		r0 = rf(tx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlaweventRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type FlaweventRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - tx *gorm.DB
//   - id uuid.UUID
func (_e *FlaweventRepository_Expecter) Delete(tx interface{}, id interface{}) *FlaweventRepository_Delete_Call {
	return &FlaweventRepository_Delete_Call{Call: _e.mock.On("Delete", tx, id)}
}

func (_c *FlaweventRepository_Delete_Call) Run(run func(tx *gorm.DB, id uuid.UUID)) *FlaweventRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *FlaweventRepository_Delete_Call) Return(_a0 error) *FlaweventRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FlaweventRepository_Delete_Call) RunAndReturn(run func(*gorm.DB, uuid.UUID) error) *FlaweventRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ids
func (_m *FlaweventRepository) List(ids []uuid.UUID) ([]flaw.Model, error) {
	ret := _m.Called(ids)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []flaw.Model
	var r1 error
	if rf, ok := ret.Get(0).(func([]uuid.UUID) ([]flaw.Model, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]uuid.UUID) []flaw.Model); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]flaw.Model)
		}
	}

	if rf, ok := ret.Get(1).(func([]uuid.UUID) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FlaweventRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type FlaweventRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ids []uuid.UUID
func (_e *FlaweventRepository_Expecter) List(ids interface{}) *FlaweventRepository_List_Call {
	return &FlaweventRepository_List_Call{Call: _e.mock.On("List", ids)}
}

func (_c *FlaweventRepository_List_Call) Run(run func(ids []uuid.UUID)) *FlaweventRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]uuid.UUID))
	})
	return _c
}

func (_c *FlaweventRepository_List_Call) Return(_a0 []flaw.Model, _a1 error) *FlaweventRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FlaweventRepository_List_Call) RunAndReturn(run func([]uuid.UUID) ([]flaw.Model, error)) *FlaweventRepository_List_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: id
func (_m *FlaweventRepository) Read(id uuid.UUID) (flaw.Model, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 flaw.Model
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (flaw.Model, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) flaw.Model); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(flaw.Model)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FlaweventRepository_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type FlaweventRepository_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - id uuid.UUID
func (_e *FlaweventRepository_Expecter) Read(id interface{}) *FlaweventRepository_Read_Call {
	return &FlaweventRepository_Read_Call{Call: _e.mock.On("Read", id)}
}

func (_c *FlaweventRepository_Read_Call) Run(run func(id uuid.UUID)) *FlaweventRepository_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *FlaweventRepository_Read_Call) Return(_a0 flaw.Model, _a1 error) *FlaweventRepository_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FlaweventRepository_Read_Call) RunAndReturn(run func(uuid.UUID) (flaw.Model, error)) *FlaweventRepository_Read_Call {
	_c.Call.Return(run)
	return _c
}

// Transaction provides a mock function with given fields: _a0
func (_m *FlaweventRepository) Transaction(_a0 func(*gorm.DB) error) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Transaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(func(*gorm.DB) error) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlaweventRepository_Transaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Transaction'
type FlaweventRepository_Transaction_Call struct {
	*mock.Call
}

// Transaction is a helper method to define mock.On call
//   - _a0 func(*gorm.DB) error
func (_e *FlaweventRepository_Expecter) Transaction(_a0 interface{}) *FlaweventRepository_Transaction_Call {
	return &FlaweventRepository_Transaction_Call{Call: _e.mock.On("Transaction", _a0)}
}

func (_c *FlaweventRepository_Transaction_Call) Run(run func(_a0 func(*gorm.DB) error)) *FlaweventRepository_Transaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(*gorm.DB) error))
	})
	return _c
}

func (_c *FlaweventRepository_Transaction_Call) Return(_a0 error) *FlaweventRepository_Transaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FlaweventRepository_Transaction_Call) RunAndReturn(run func(func(*gorm.DB) error) error) *FlaweventRepository_Transaction_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: tx, t
func (_m *FlaweventRepository) Update(tx *gorm.DB, t *flaw.Model) error {
	ret := _m.Called(tx, t)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *flaw.Model) error); ok {
		r0 = rf(tx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlaweventRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type FlaweventRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - tx *gorm.DB
//   - t *flaw.Model
func (_e *FlaweventRepository_Expecter) Update(tx interface{}, t interface{}) *FlaweventRepository_Update_Call {
	return &FlaweventRepository_Update_Call{Call: _e.mock.On("Update", tx, t)}
}

func (_c *FlaweventRepository_Update_Call) Run(run func(tx *gorm.DB, t *flaw.Model)) *FlaweventRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*flaw.Model))
	})
	return _c
}

func (_c *FlaweventRepository_Update_Call) Return(_a0 error) *FlaweventRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FlaweventRepository_Update_Call) RunAndReturn(run func(*gorm.DB, *flaw.Model) error) *FlaweventRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewFlaweventRepository creates a new instance of FlaweventRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFlaweventRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *FlaweventRepository {
	mock := &FlaweventRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

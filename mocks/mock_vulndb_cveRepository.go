// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	models "github.com/l3montree-dev/flawfix/internal/database/models"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"

	time "time"
)

// VulndbCveRepository is an autogenerated mock type for the cveRepository type
type VulndbCveRepository struct {
	mock.Mock
}

type VulndbCveRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *VulndbCveRepository) EXPECT() *VulndbCveRepository_Expecter {
	return &VulndbCveRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: tx, t
func (_m *VulndbCveRepository) Create(tx *gorm.DB, t *models.CVE) error {
	ret := _m.Called(tx, t)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.CVE) error); ok {
		r0 = rf(tx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VulndbCveRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type VulndbCveRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - tx *gorm.DB
//   - t *models.CVE
func (_e *VulndbCveRepository_Expecter) Create(tx interface{}, t interface{}) *VulndbCveRepository_Create_Call {
	return &VulndbCveRepository_Create_Call{Call: _e.mock.On("Create", tx, t)}
}

func (_c *VulndbCveRepository_Create_Call) Run(run func(tx *gorm.DB, t *models.CVE)) *VulndbCveRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*models.CVE))
	})
	return _c
}

func (_c *VulndbCveRepository_Create_Call) Return(_a0 error) *VulndbCveRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *VulndbCveRepository_Create_Call) RunAndReturn(run func(*gorm.DB, *models.CVE) error) *VulndbCveRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateBatch provides a mock function with given fields: tx, ts
func (_m *VulndbCveRepository) CreateBatch(tx *gorm.DB, ts []models.CVE) error {
	ret := _m.Called(tx, ts)

	if len(ret) == 0 {
		panic("no return value specified for CreateBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, []models.CVE) error); ok {
		r0 = rf(tx, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VulndbCveRepository_CreateBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBatch'
type VulndbCveRepository_CreateBatch_Call struct {
	*mock.Call
}

// CreateBatch is a helper method to define mock.On call
//   - tx *gorm.DB
//   - ts []models.CVE
func (_e *VulndbCveRepository_Expecter) CreateBatch(tx interface{}, ts interface{}) *VulndbCveRepository_CreateBatch_Call {
	return &VulndbCveRepository_CreateBatch_Call{Call: _e.mock.On("CreateBatch", tx, ts)}
}

func (_c *VulndbCveRepository_CreateBatch_Call) Run(run func(tx *gorm.DB, ts []models.CVE)) *VulndbCveRepository_CreateBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].([]models.CVE))
	})
	return _c
}

func (_c *VulndbCveRepository_CreateBatch_Call) Return(_a0 error) *VulndbCveRepository_CreateBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *VulndbCveRepository_CreateBatch_Call) RunAndReturn(run func(*gorm.DB, []models.CVE) error) *VulndbCveRepository_CreateBatch_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: tx, id
func (_m *VulndbCveRepository) Delete(tx *gorm.DB, id string) error {
	ret := _m.Called(tx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) error); ok {
		r0 = rf(tx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VulndbCveRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type VulndbCveRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - tx *gorm.DB
//   - id string
func (_e *VulndbCveRepository_Expecter) Delete(tx interface{}, id interface{}) *VulndbCveRepository_Delete_Call {
	return &VulndbCveRepository_Delete_Call{Call: _e.mock.On("Delete", tx, id)}
}

func (_c *VulndbCveRepository_Delete_Call) Run(run func(tx *gorm.DB, id string)) *VulndbCveRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(string))
	})
	return _c
}

func (_c *VulndbCveRepository_Delete_Call) Return(_a0 error) *VulndbCveRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *VulndbCveRepository_Delete_Call) RunAndReturn(run func(*gorm.DB, string) error) *VulndbCveRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// FindByID provides a mock function with given fields: id
func (_m *VulndbCveRepository) FindByID(id string) (models.CVE, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 models.CVE
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.CVE, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.CVE); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.CVE)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VulndbCveRepository_FindByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByID'
type VulndbCveRepository_FindByID_Call struct {
	*mock.Call
}

// FindByID is a helper method to define mock.On call
//   - id string
func (_e *VulndbCveRepository_Expecter) FindByID(id interface{}) *VulndbCveRepository_FindByID_Call {
	return &VulndbCveRepository_FindByID_Call{Call: _e.mock.On("FindByID", id)}
}

func (_c *VulndbCveRepository_FindByID_Call) Run(run func(id string)) *VulndbCveRepository_FindByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *VulndbCveRepository_FindByID_Call) Return(_a0 models.CVE, _a1 error) *VulndbCveRepository_FindByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VulndbCveRepository_FindByID_Call) RunAndReturn(run func(string) (models.CVE, error)) *VulndbCveRepository_FindByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetDB provides a mock function with given fields: tx
func (_m *VulndbCveRepository) GetDB(tx *gorm.DB) *gorm.DB {
	ret := _m.Called(tx)

	if len(ret) == 0 {
		panic("no return value specified for GetDB")
	}

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(*gorm.DB) *gorm.DB); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// VulndbCveRepository_GetDB_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDB'
type VulndbCveRepository_GetDB_Call struct {
	*mock.Call
}

// GetDB is a helper method to define mock.On call
//   - tx *gorm.DB
func (_e *VulndbCveRepository_Expecter) GetDB(tx interface{}) *VulndbCveRepository_GetDB_Call {
	return &VulndbCveRepository_GetDB_Call{Call: _e.mock.On("GetDB", tx)}
}

func (_c *VulndbCveRepository_GetDB_Call) Run(run func(tx *gorm.DB)) *VulndbCveRepository_GetDB_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB))
	})
	return _c
}

func (_c *VulndbCveRepository_GetDB_Call) Return(_a0 *gorm.DB) *VulndbCveRepository_GetDB_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *VulndbCveRepository_GetDB_Call) RunAndReturn(run func(*gorm.DB) *gorm.DB) *VulndbCveRepository_GetDB_Call {
	_c.Call.Return(run)
	return _c
}

// GetLastModDate provides a mock function with given fields:
func (_m *VulndbCveRepository) GetLastModDate() (time.Time, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLastModDate")
	}

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func() (time.Time, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VulndbCveRepository_GetLastModDate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastModDate'
type VulndbCveRepository_GetLastModDate_Call struct {
	*mock.Call
}

// GetLastModDate is a helper method to define mock.On call
func (_e *VulndbCveRepository_Expecter) GetLastModDate() *VulndbCveRepository_GetLastModDate_Call {
	return &VulndbCveRepository_GetLastModDate_Call{Call: _e.mock.On("GetLastModDate")}
}

func (_c *VulndbCveRepository_GetLastModDate_Call) Run(run func()) *VulndbCveRepository_GetLastModDate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *VulndbCveRepository_GetLastModDate_Call) Return(_a0 time.Time, _a1 error) *VulndbCveRepository_GetLastModDate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VulndbCveRepository_GetLastModDate_Call) RunAndReturn(run func() (time.Time, error)) *VulndbCveRepository_GetLastModDate_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ids
func (_m *VulndbCveRepository) List(ids []string) ([]models.CVE, error) {
	ret := _m.Called(ids)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []models.CVE
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]models.CVE, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]string) []models.CVE); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.CVE)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VulndbCveRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type VulndbCveRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ids []string
func (_e *VulndbCveRepository_Expecter) List(ids interface{}) *VulndbCveRepository_List_Call {
	return &VulndbCveRepository_List_Call{Call: _e.mock.On("List", ids)}
}

func (_c *VulndbCveRepository_List_Call) Run(run func(ids []string)) *VulndbCveRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *VulndbCveRepository_List_Call) Return(_a0 []models.CVE, _a1 error) *VulndbCveRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VulndbCveRepository_List_Call) RunAndReturn(run func([]string) ([]models.CVE, error)) *VulndbCveRepository_List_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: id
func (_m *VulndbCveRepository) Read(id string) (models.CVE, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 models.CVE
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.CVE, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.CVE); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.CVE)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VulndbCveRepository_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type VulndbCveRepository_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - id string
func (_e *VulndbCveRepository_Expecter) Read(id interface{}) *VulndbCveRepository_Read_Call {
	return &VulndbCveRepository_Read_Call{Call: _e.mock.On("Read", id)}
}

func (_c *VulndbCveRepository_Read_Call) Run(run func(id string)) *VulndbCveRepository_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *VulndbCveRepository_Read_Call) Return(_a0 models.CVE, _a1 error) *VulndbCveRepository_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VulndbCveRepository_Read_Call) RunAndReturn(run func(string) (models.CVE, error)) *VulndbCveRepository_Read_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: tx, t
func (_m *VulndbCveRepository) Save(tx *gorm.DB, t *models.CVE) error {
	ret := _m.Called(tx, t)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.CVE) error); ok {
		r0 = rf(tx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VulndbCveRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type VulndbCveRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - tx *gorm.DB
//   - t *models.CVE
func (_e *VulndbCveRepository_Expecter) Save(tx interface{}, t interface{}) *VulndbCveRepository_Save_Call {
	return &VulndbCveRepository_Save_Call{Call: _e.mock.On("Save", tx, t)}
}

func (_c *VulndbCveRepository_Save_Call) Run(run func(tx *gorm.DB, t *models.CVE)) *VulndbCveRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*models.CVE))
	})
	return _c
}

func (_c *VulndbCveRepository_Save_Call) Return(_a0 error) *VulndbCveRepository_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *VulndbCveRepository_Save_Call) RunAndReturn(run func(*gorm.DB, *models.CVE) error) *VulndbCveRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// SaveBatch provides a mock function with given fields: tx, ts
func (_m *VulndbCveRepository) SaveBatch(tx *gorm.DB, ts []models.CVE) error {
	ret := _m.Called(tx, ts)

	if len(ret) == 0 {
		panic("no return value specified for SaveBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, []models.CVE) error); ok {
		r0 = rf(tx, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VulndbCveRepository_SaveBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveBatch'
type VulndbCveRepository_SaveBatch_Call struct {
	*mock.Call
}

// SaveBatch is a helper method to define mock.On call
//   - tx *gorm.DB
//   - ts []models.CVE
func (_e *VulndbCveRepository_Expecter) SaveBatch(tx interface{}, ts interface{}) *VulndbCveRepository_SaveBatch_Call {
	return &VulndbCveRepository_SaveBatch_Call{Call: _e.mock.On("SaveBatch", tx, ts)}
}

func (_c *VulndbCveRepository_SaveBatch_Call) Run(run func(tx *gorm.DB, ts []models.CVE)) *VulndbCveRepository_SaveBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].([]models.CVE))
	})
	return _c
}

func (_c *VulndbCveRepository_SaveBatch_Call) Return(_a0 error) *VulndbCveRepository_SaveBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *VulndbCveRepository_SaveBatch_Call) RunAndReturn(run func(*gorm.DB, []models.CVE) error) *VulndbCveRepository_SaveBatch_Call {
	_c.Call.Return(run)
	return _c
}

// Transaction provides a mock function with given fields: _a0
func (_m *VulndbCveRepository) Transaction(_a0 func(*gorm.DB) error) error {
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

// VulndbCveRepository_Transaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Transaction'
type VulndbCveRepository_Transaction_Call struct {
	*mock.Call
}

// Transaction is a helper method to define mock.On call
//   - _a0 func(*gorm.DB) error
func (_e *VulndbCveRepository_Expecter) Transaction(_a0 interface{}) *VulndbCveRepository_Transaction_Call {
	return &VulndbCveRepository_Transaction_Call{Call: _e.mock.On("Transaction", _a0)}
}

func (_c *VulndbCveRepository_Transaction_Call) Run(run func(_a0 func(*gorm.DB) error)) *VulndbCveRepository_Transaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(*gorm.DB) error))
	})
	return _c
}

func (_c *VulndbCveRepository_Transaction_Call) Return(_a0 error) *VulndbCveRepository_Transaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *VulndbCveRepository_Transaction_Call) RunAndReturn(run func(func(*gorm.DB) error) error) *VulndbCveRepository_Transaction_Call {
	_c.Call.Return(run)
	return _c
}

// NewVulndbCveRepository creates a new instance of VulndbCveRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVulndbCveRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *VulndbCveRepository {
	mock := &VulndbCveRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

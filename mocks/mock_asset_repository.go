// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	models "github.com/l3montree-dev/flawfix/internal/database/models"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"

	uuid "github.com/google/uuid"
)

// AssetRepository is an autogenerated mock type for the repository type
type AssetRepository struct {
	mock.Mock
}

type AssetRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *AssetRepository) EXPECT() *AssetRepository_Expecter {
	return &AssetRepository_Expecter{mock: &_m.Mock}
}

// Begin provides a mock function with given fields:
func (_m *AssetRepository) Begin() *gorm.DB {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Begin")
	}

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// AssetRepository_Begin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Begin'
type AssetRepository_Begin_Call struct {
	*mock.Call
}

// Begin is a helper method to define mock.On call
func (_e *AssetRepository_Expecter) Begin() *AssetRepository_Begin_Call {
	return &AssetRepository_Begin_Call{Call: _e.mock.On("Begin")}
}

func (_c *AssetRepository_Begin_Call) Run(run func()) *AssetRepository_Begin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AssetRepository_Begin_Call) Return(_a0 *gorm.DB) *AssetRepository_Begin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AssetRepository_Begin_Call) RunAndReturn(run func() *gorm.DB) *AssetRepository_Begin_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: tx, t
func (_m *AssetRepository) Create(tx *gorm.DB, t *models.Asset) error {
	ret := _m.Called(tx, t)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.Asset) error); ok {
		r0 = rf(tx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AssetRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type AssetRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - tx *gorm.DB
//   - t *models.Asset
func (_e *AssetRepository_Expecter) Create(tx interface{}, t interface{}) *AssetRepository_Create_Call {
	return &AssetRepository_Create_Call{Call: _e.mock.On("Create", tx, t)}
}

func (_c *AssetRepository_Create_Call) Run(run func(tx *gorm.DB, t *models.Asset)) *AssetRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*models.Asset))
	})
	return _c
}

func (_c *AssetRepository_Create_Call) Return(_a0 error) *AssetRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AssetRepository_Create_Call) RunAndReturn(run func(*gorm.DB, *models.Asset) error) *AssetRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateBatch provides a mock function with given fields: tx, ts
func (_m *AssetRepository) CreateBatch(tx *gorm.DB, ts []models.Asset) error {
	ret := _m.Called(tx, ts)

	if len(ret) == 0 {
		panic("no return value specified for CreateBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, []models.Asset) error); ok {
		r0 = rf(tx, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AssetRepository_CreateBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBatch'
type AssetRepository_CreateBatch_Call struct {
	*mock.Call
}

// CreateBatch is a helper method to define mock.On call
//   - tx *gorm.DB
//   - ts []models.Asset
func (_e *AssetRepository_Expecter) CreateBatch(tx interface{}, ts interface{}) *AssetRepository_CreateBatch_Call {
	return &AssetRepository_CreateBatch_Call{Call: _e.mock.On("CreateBatch", tx, ts)}
}

func (_c *AssetRepository_CreateBatch_Call) Run(run func(tx *gorm.DB, ts []models.Asset)) *AssetRepository_CreateBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].([]models.Asset))
	})
	return _c
}

func (_c *AssetRepository_CreateBatch_Call) Return(_a0 error) *AssetRepository_CreateBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AssetRepository_CreateBatch_Call) RunAndReturn(run func(*gorm.DB, []models.Asset) error) *AssetRepository_CreateBatch_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: tx, id
func (_m *AssetRepository) Delete(tx *gorm.DB, id uuid.UUID) error {
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

// AssetRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type AssetRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - tx *gorm.DB
//   - id uuid.UUID
func (_e *AssetRepository_Expecter) Delete(tx interface{}, id interface{}) *AssetRepository_Delete_Call {
	return &AssetRepository_Delete_Call{Call: _e.mock.On("Delete", tx, id)}
}

func (_c *AssetRepository_Delete_Call) Run(run func(tx *gorm.DB, id uuid.UUID)) *AssetRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *AssetRepository_Delete_Call) Return(_a0 error) *AssetRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AssetRepository_Delete_Call) RunAndReturn(run func(*gorm.DB, uuid.UUID) error) *AssetRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// FindByName provides a mock function with given fields: name
func (_m *AssetRepository) FindByName(name string) (models.Asset, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for FindByName")
	}

	var r0 models.Asset
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Asset, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) models.Asset); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.Asset)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetRepository_FindByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByName'
type AssetRepository_FindByName_Call struct {
	*mock.Call
}

// FindByName is a helper method to define mock.On call
//   - name string
func (_e *AssetRepository_Expecter) FindByName(name interface{}) *AssetRepository_FindByName_Call {
	return &AssetRepository_FindByName_Call{Call: _e.mock.On("FindByName", name)}
}

func (_c *AssetRepository_FindByName_Call) Run(run func(name string)) *AssetRepository_FindByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AssetRepository_FindByName_Call) Return(_a0 models.Asset, _a1 error) *AssetRepository_FindByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AssetRepository_FindByName_Call) RunAndReturn(run func(string) (models.Asset, error)) *AssetRepository_FindByName_Call {
	_c.Call.Return(run)
	return _c
}

// FindOrCreate provides a mock function with given fields: tx, name
func (_m *AssetRepository) FindOrCreate(tx *gorm.DB, name string) (models.Asset, error) {
	ret := _m.Called(tx, name)

	if len(ret) == 0 {
		panic("no return value specified for FindOrCreate")
	}

	var r0 models.Asset
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) (models.Asset, error)); ok {
		return rf(tx, name)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) models.Asset); ok {
		r0 = rf(tx, name)
	} else {
		r0 = ret.Get(0).(models.Asset)
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, string) error); ok {
		r1 = rf(tx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetRepository_FindOrCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOrCreate'
type AssetRepository_FindOrCreate_Call struct {
	*mock.Call
}

// FindOrCreate is a helper method to define mock.On call
//   - tx *gorm.DB
//   - name string
func (_e *AssetRepository_Expecter) FindOrCreate(tx interface{}, name interface{}) *AssetRepository_FindOrCreate_Call {
	return &AssetRepository_FindOrCreate_Call{Call: _e.mock.On("FindOrCreate", tx, name)}
}

func (_c *AssetRepository_FindOrCreate_Call) Run(run func(tx *gorm.DB, name string)) *AssetRepository_FindOrCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(string))
	})
	return _c
}

func (_c *AssetRepository_FindOrCreate_Call) Return(_a0 models.Asset, _a1 error) *AssetRepository_FindOrCreate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AssetRepository_FindOrCreate_Call) RunAndReturn(run func(*gorm.DB, string) (models.Asset, error)) *AssetRepository_FindOrCreate_Call {
	_c.Call.Return(run)
	return _c
}

// GetAssetIDBySlug provides a mock function with given fields: projectID, slug
func (_m *AssetRepository) GetAssetIDBySlug(projectID uuid.UUID, slug string) (uuid.UUID, error) {
	ret := _m.Called(projectID, slug)

	if len(ret) == 0 {
		panic("no return value specified for GetAssetIDBySlug")
	}

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, string) (uuid.UUID, error)); ok {
		return rf(projectID, slug)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID, string) uuid.UUID); ok {
		r0 = rf(projectID, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID, string) error); ok {
		r1 = rf(projectID, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetRepository_GetAssetIDBySlug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAssetIDBySlug'
type AssetRepository_GetAssetIDBySlug_Call struct {
	*mock.Call
}

// GetAssetIDBySlug is a helper method to define mock.On call
//   - projectID uuid.UUID
//   - slug string
func (_e *AssetRepository_Expecter) GetAssetIDBySlug(projectID interface{}, slug interface{}) *AssetRepository_GetAssetIDBySlug_Call {
	return &AssetRepository_GetAssetIDBySlug_Call{Call: _e.mock.On("GetAssetIDBySlug", projectID, slug)}
}

func (_c *AssetRepository_GetAssetIDBySlug_Call) Run(run func(projectID uuid.UUID, slug string)) *AssetRepository_GetAssetIDBySlug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID), args[1].(string))
	})
	return _c
}

func (_c *AssetRepository_GetAssetIDBySlug_Call) Return(_a0 uuid.UUID, _a1 error) *AssetRepository_GetAssetIDBySlug_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AssetRepository_GetAssetIDBySlug_Call) RunAndReturn(run func(uuid.UUID, string) (uuid.UUID, error)) *AssetRepository_GetAssetIDBySlug_Call {
	_c.Call.Return(run)
	return _c
}

// GetByProjectID provides a mock function with given fields: projectID
func (_m *AssetRepository) GetByProjectID(projectID uuid.UUID) ([]models.Asset, error) {
	ret := _m.Called(projectID)

	if len(ret) == 0 {
		panic("no return value specified for GetByProjectID")
	}

	var r0 []models.Asset
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) ([]models.Asset, error)); ok {
		return rf(projectID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) []models.Asset); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Asset)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetRepository_GetByProjectID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByProjectID'
type AssetRepository_GetByProjectID_Call struct {
	*mock.Call
}

// GetByProjectID is a helper method to define mock.On call
//   - projectID uuid.UUID
func (_e *AssetRepository_Expecter) GetByProjectID(projectID interface{}) *AssetRepository_GetByProjectID_Call {
	return &AssetRepository_GetByProjectID_Call{Call: _e.mock.On("GetByProjectID", projectID)}
}

func (_c *AssetRepository_GetByProjectID_Call) Run(run func(projectID uuid.UUID)) *AssetRepository_GetByProjectID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *AssetRepository_GetByProjectID_Call) Return(_a0 []models.Asset, _a1 error) *AssetRepository_GetByProjectID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AssetRepository_GetByProjectID_Call) RunAndReturn(run func(uuid.UUID) ([]models.Asset, error)) *AssetRepository_GetByProjectID_Call {
	_c.Call.Return(run)
	return _c
}

// GetDB provides a mock function with given fields: tx
func (_m *AssetRepository) GetDB(tx *gorm.DB) *gorm.DB {
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

// AssetRepository_GetDB_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDB'
type AssetRepository_GetDB_Call struct {
	*mock.Call
}

// GetDB is a helper method to define mock.On call
//   - tx *gorm.DB
func (_e *AssetRepository_Expecter) GetDB(tx interface{}) *AssetRepository_GetDB_Call {
	return &AssetRepository_GetDB_Call{Call: _e.mock.On("GetDB", tx)}
}

func (_c *AssetRepository_GetDB_Call) Run(run func(tx *gorm.DB)) *AssetRepository_GetDB_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB))
	})
	return _c
}

func (_c *AssetRepository_GetDB_Call) Return(_a0 *gorm.DB) *AssetRepository_GetDB_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AssetRepository_GetDB_Call) RunAndReturn(run func(*gorm.DB) *gorm.DB) *AssetRepository_GetDB_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ids
func (_m *AssetRepository) List(ids []uuid.UUID) ([]models.Asset, error) {
	ret := _m.Called(ids)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []models.Asset
	var r1 error
	if rf, ok := ret.Get(0).(func([]uuid.UUID) ([]models.Asset, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]uuid.UUID) []models.Asset); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Asset)
		}
	}

	if rf, ok := ret.Get(1).(func([]uuid.UUID) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type AssetRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ids []uuid.UUID
func (_e *AssetRepository_Expecter) List(ids interface{}) *AssetRepository_List_Call {
	return &AssetRepository_List_Call{Call: _e.mock.On("List", ids)}
}

func (_c *AssetRepository_List_Call) Run(run func(ids []uuid.UUID)) *AssetRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]uuid.UUID))
	})
	return _c
}

func (_c *AssetRepository_List_Call) Return(_a0 []models.Asset, _a1 error) *AssetRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AssetRepository_List_Call) RunAndReturn(run func([]uuid.UUID) ([]models.Asset, error)) *AssetRepository_List_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: id
func (_m *AssetRepository) Read(id uuid.UUID) (models.Asset, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 models.Asset
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (models.Asset, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) models.Asset); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Asset)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetRepository_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type AssetRepository_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - id uuid.UUID
func (_e *AssetRepository_Expecter) Read(id interface{}) *AssetRepository_Read_Call {
	return &AssetRepository_Read_Call{Call: _e.mock.On("Read", id)}
}

func (_c *AssetRepository_Read_Call) Run(run func(id uuid.UUID)) *AssetRepository_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *AssetRepository_Read_Call) Return(_a0 models.Asset, _a1 error) *AssetRepository_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AssetRepository_Read_Call) RunAndReturn(run func(uuid.UUID) (models.Asset, error)) *AssetRepository_Read_Call {
	_c.Call.Return(run)
	return _c
}

// ReadBySlug provides a mock function with given fields: projectID, slug
func (_m *AssetRepository) ReadBySlug(projectID uuid.UUID, slug string) (models.Asset, error) {
	ret := _m.Called(projectID, slug)

	if len(ret) == 0 {
		panic("no return value specified for ReadBySlug")
	}

	var r0 models.Asset
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, string) (models.Asset, error)); ok {
		return rf(projectID, slug)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID, string) models.Asset); ok {
		r0 = rf(projectID, slug)
	} else {
		r0 = ret.Get(0).(models.Asset)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID, string) error); ok {
		r1 = rf(projectID, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetRepository_ReadBySlug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadBySlug'
type AssetRepository_ReadBySlug_Call struct {
	*mock.Call
}

// ReadBySlug is a helper method to define mock.On call
//   - projectID uuid.UUID
//   - slug string
func (_e *AssetRepository_Expecter) ReadBySlug(projectID interface{}, slug interface{}) *AssetRepository_ReadBySlug_Call {
	return &AssetRepository_ReadBySlug_Call{Call: _e.mock.On("ReadBySlug", projectID, slug)}
}

func (_c *AssetRepository_ReadBySlug_Call) Run(run func(projectID uuid.UUID, slug string)) *AssetRepository_ReadBySlug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID), args[1].(string))
	})
	return _c
}

func (_c *AssetRepository_ReadBySlug_Call) Return(_a0 models.Asset, _a1 error) *AssetRepository_ReadBySlug_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AssetRepository_ReadBySlug_Call) RunAndReturn(run func(uuid.UUID, string) (models.Asset, error)) *AssetRepository_ReadBySlug_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: tx, t
func (_m *AssetRepository) Save(tx *gorm.DB, t *models.Asset) error {
	ret := _m.Called(tx, t)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.Asset) error); ok {
		r0 = rf(tx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AssetRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type AssetRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - tx *gorm.DB
//   - t *models.Asset
func (_e *AssetRepository_Expecter) Save(tx interface{}, t interface{}) *AssetRepository_Save_Call {
	return &AssetRepository_Save_Call{Call: _e.mock.On("Save", tx, t)}
}

func (_c *AssetRepository_Save_Call) Run(run func(tx *gorm.DB, t *models.Asset)) *AssetRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*models.Asset))
	})
	return _c
}

func (_c *AssetRepository_Save_Call) Return(_a0 error) *AssetRepository_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AssetRepository_Save_Call) RunAndReturn(run func(*gorm.DB, *models.Asset) error) *AssetRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// SaveBatch provides a mock function with given fields: tx, ts
func (_m *AssetRepository) SaveBatch(tx *gorm.DB, ts []models.Asset) error {
	ret := _m.Called(tx, ts)

	if len(ret) == 0 {
		panic("no return value specified for SaveBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, []models.Asset) error); ok {
		r0 = rf(tx, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AssetRepository_SaveBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveBatch'
type AssetRepository_SaveBatch_Call struct {
	*mock.Call
}

// SaveBatch is a helper method to define mock.On call
//   - tx *gorm.DB
//   - ts []models.Asset
func (_e *AssetRepository_Expecter) SaveBatch(tx interface{}, ts interface{}) *AssetRepository_SaveBatch_Call {
	return &AssetRepository_SaveBatch_Call{Call: _e.mock.On("SaveBatch", tx, ts)}
}

func (_c *AssetRepository_SaveBatch_Call) Run(run func(tx *gorm.DB, ts []models.Asset)) *AssetRepository_SaveBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].([]models.Asset))
	})
	return _c
}

func (_c *AssetRepository_SaveBatch_Call) Return(_a0 error) *AssetRepository_SaveBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AssetRepository_SaveBatch_Call) RunAndReturn(run func(*gorm.DB, []models.Asset) error) *AssetRepository_SaveBatch_Call {
	_c.Call.Return(run)
	return _c
}

// Transaction provides a mock function with given fields: _a0
func (_m *AssetRepository) Transaction(_a0 func(*gorm.DB) error) error {
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

// AssetRepository_Transaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Transaction'
type AssetRepository_Transaction_Call struct {
	*mock.Call
}

// Transaction is a helper method to define mock.On call
//   - _a0 func(*gorm.DB) error
func (_e *AssetRepository_Expecter) Transaction(_a0 interface{}) *AssetRepository_Transaction_Call {
	return &AssetRepository_Transaction_Call{Call: _e.mock.On("Transaction", _a0)}
}

func (_c *AssetRepository_Transaction_Call) Run(run func(_a0 func(*gorm.DB) error)) *AssetRepository_Transaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(*gorm.DB) error))
	})
	return _c
}

func (_c *AssetRepository_Transaction_Call) Return(_a0 error) *AssetRepository_Transaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AssetRepository_Transaction_Call) RunAndReturn(run func(func(*gorm.DB) error) error) *AssetRepository_Transaction_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: tx, _a1
func (_m *AssetRepository) Update(tx *gorm.DB, _a1 *models.Asset) error {
	ret := _m.Called(tx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.Asset) error); ok {
		r0 = rf(tx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AssetRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type AssetRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - tx *gorm.DB
//   - _a1 *models.Asset
func (_e *AssetRepository_Expecter) Update(tx interface{}, _a1 interface{}) *AssetRepository_Update_Call {
	return &AssetRepository_Update_Call{Call: _e.mock.On("Update", tx, _a1)}
}

func (_c *AssetRepository_Update_Call) Run(run func(tx *gorm.DB, _a1 *models.Asset)) *AssetRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*models.Asset))
	})
	return _c
}

func (_c *AssetRepository_Update_Call) Return(_a0 error) *AssetRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AssetRepository_Update_Call) RunAndReturn(run func(*gorm.DB, *models.Asset) error) *AssetRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewAssetRepository creates a new instance of AssetRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAssetRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AssetRepository {
	mock := &AssetRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

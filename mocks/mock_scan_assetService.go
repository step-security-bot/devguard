// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	cyclonedx "github.com/CycloneDX/cyclonedx-go"
	mock "github.com/stretchr/testify/mock"

	models "github.com/l3montree-dev/devguard/internal/database/models"
)

// ScanAssetService is an autogenerated mock type for the assetService type
type ScanAssetService struct {
	mock.Mock
}

type ScanAssetService_Expecter struct {
	mock *mock.Mock
}

func (_m *ScanAssetService) EXPECT() *ScanAssetService_Expecter {
	return &ScanAssetService_Expecter{mock: &_m.Mock}
}

// HandleScanResult provides a mock function with given fields: user, scannerID, asset, flaws
func (_m *ScanAssetService) HandleScanResult(user string, scannerID string, asset models.Asset, flaws []models.Flaw) (int, int, []models.Flaw, error) {
	ret := _m.Called(user, scannerID, asset, flaws)

	if len(ret) == 0 {
		panic("no return value specified for HandleScanResult")
	}

	var r0 int
	var r1 int
	var r2 []models.Flaw
	var r3 error
	if rf, ok := ret.Get(0).(func(string, string, models.Asset, []models.Flaw) (int, int, []models.Flaw, error)); ok {
		return rf(user, scannerID, asset, flaws)
	}
	if rf, ok := ret.Get(0).(func(string, string, models.Asset, []models.Flaw) int); ok {
		r0 = rf(user, scannerID, asset, flaws)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string, string, models.Asset, []models.Flaw) int); ok {
		r1 = rf(user, scannerID, asset, flaws)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(string, string, models.Asset, []models.Flaw) []models.Flaw); ok {
		r2 = rf(user, scannerID, asset, flaws)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]models.Flaw)
		}
	}

	if rf, ok := ret.Get(3).(func(string, string, models.Asset, []models.Flaw) error); ok {
		r3 = rf(user, scannerID, asset, flaws)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// ScanAssetService_HandleScanResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleScanResult'
type ScanAssetService_HandleScanResult_Call struct {
	*mock.Call
}

// HandleScanResult is a helper method to define mock.On call
//   - user string
//   - scannerID string
//   - asset models.Asset
//   - flaws []models.Flaw
func (_e *ScanAssetService_Expecter) HandleScanResult(user interface{}, scannerID interface{}, asset interface{}, flaws interface{}) *ScanAssetService_HandleScanResult_Call {
	return &ScanAssetService_HandleScanResult_Call{Call: _e.mock.On("HandleScanResult", user, scannerID, asset, flaws)}
}

func (_c *ScanAssetService_HandleScanResult_Call) Run(run func(user string, scannerID string, asset models.Asset, flaws []models.Flaw)) *ScanAssetService_HandleScanResult_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(models.Asset), args[3].([]models.Flaw))
	})
	return _c
}

func (_c *ScanAssetService_HandleScanResult_Call) Return(amountOpened int, amountClosed int, newState []models.Flaw, err error) *ScanAssetService_HandleScanResult_Call {
	_c.Call.Return(amountOpened, amountClosed, newState, err)
	return _c
}

func (_c *ScanAssetService_HandleScanResult_Call) RunAndReturn(run func(string, string, models.Asset, []models.Flaw) (int, int, []models.Flaw, error)) *ScanAssetService_HandleScanResult_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSBOM provides a mock function with given fields: asset, scanType, version, sbom
func (_m *ScanAssetService) UpdateSBOM(asset models.Asset, scanType string, version string, sbom *cyclonedx.BOM) error {
	ret := _m.Called(asset, scanType, version, sbom)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSBOM")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Asset, string, string, *cyclonedx.BOM) error); ok {
		r0 = rf(asset, scanType, version, sbom)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ScanAssetService_UpdateSBOM_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSBOM'
type ScanAssetService_UpdateSBOM_Call struct {
	*mock.Call
}

// UpdateSBOM is a helper method to define mock.On call
//   - asset models.Asset
//   - scanType string
//   - version string
//   - sbom *cyclonedx.BOM
func (_e *ScanAssetService_Expecter) UpdateSBOM(asset interface{}, scanType interface{}, version interface{}, sbom interface{}) *ScanAssetService_UpdateSBOM_Call {
	return &ScanAssetService_UpdateSBOM_Call{Call: _e.mock.On("UpdateSBOM", asset, scanType, version, sbom)}
}

func (_c *ScanAssetService_UpdateSBOM_Call) Run(run func(asset models.Asset, scanType string, version string, sbom *cyclonedx.BOM)) *ScanAssetService_UpdateSBOM_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.Asset), args[1].(string), args[2].(string), args[3].(*cyclonedx.BOM))
	})
	return _c
}

func (_c *ScanAssetService_UpdateSBOM_Call) Return(_a0 error) *ScanAssetService_UpdateSBOM_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ScanAssetService_UpdateSBOM_Call) RunAndReturn(run func(models.Asset, string, string, *cyclonedx.BOM) error) *ScanAssetService_UpdateSBOM_Call {
	_c.Call.Return(run)
	return _c
}

// NewScanAssetService creates a new instance of ScanAssetService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewScanAssetService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ScanAssetService {
	mock := &ScanAssetService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

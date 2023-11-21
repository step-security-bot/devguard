// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// applicationService is an autogenerated mock type for the applicationService type
type applicationService struct {
	mock.Mock
}

// GetApplicationIDBySlug provides a mock function with given fields: projectID, slug
func (_m *applicationService) GetApplicationIDBySlug(projectID uuid.UUID, slug string) (uuid.UUID, error) {
	ret := _m.Called(projectID, slug)

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

// newApplicationService creates a new instance of applicationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newApplicationService(t interface {
	mock.TestingT
	Cleanup(func())
}) *applicationService {
	mock := &applicationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
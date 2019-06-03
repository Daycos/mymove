// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/transcom/mymove/pkg/models"
import services "github.com/transcom/mymove/pkg/services"

// OfficeUserListFetcher is an autogenerated mock type for the OfficeUserListFetcher type
type OfficeUserListFetcher struct {
	mock.Mock
}

// FetchOfficeUserList provides a mock function with given fields: filters
func (_m *OfficeUserListFetcher) FetchOfficeUserList(filters []services.QueryFilter) (models.OfficeUsers, error) {
	ret := _m.Called(filters)

	var r0 models.OfficeUsers
	if rf, ok := ret.Get(0).(func([]services.QueryFilter) models.OfficeUsers); ok {
		r0 = rf(filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.OfficeUsers)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]services.QueryFilter) error); ok {
		r1 = rf(filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

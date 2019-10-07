// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	services "github.com/transcom/mymove/pkg/services"
)

// ElectronicOrderListFetcher is an autogenerated mock type for the ElectronicOrderListFetcher type
type ElectronicOrderListFetcher struct {
	mock.Mock
}

// FetchElectronicOrderList provides a mock function with given fields: filters, associations, pagination
func (_m *ElectronicOrderListFetcher) FetchElectronicOrderList(filters []services.QueryFilter, associations services.QueryAssociations, pagination services.Pagination) (models.ElectronicOrders, error) {
	ret := _m.Called(filters, associations, pagination)

	var r0 models.ElectronicOrders
	if rf, ok := ret.Get(0).(func([]services.QueryFilter, services.QueryAssociations, services.Pagination) models.ElectronicOrders); ok {
		r0 = rf(filters, associations, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.ElectronicOrders)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]services.QueryFilter, services.QueryAssociations, services.Pagination) error); ok {
		r1 = rf(filters, associations, pagination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	services "github.com/transcom/mymove/pkg/services"
)

// TransportationServiceProviderPerformanceListFetcher is an autogenerated mock type for the TransportationServiceProviderPerformanceListFetcher type
type TransportationServiceProviderPerformanceListFetcher struct {
	mock.Mock
}

// FetchTransportationServiceProviderPerformanceList provides a mock function with given fields: filters, associations, pagination, ordering
func (_m *TransportationServiceProviderPerformanceListFetcher) FetchTransportationServiceProviderPerformanceList(filters []services.QueryFilter, associations services.QueryAssociations, pagination services.Pagination, ordering services.QueryOrder) (models.TransportationServiceProviderPerformances, error) {
	ret := _m.Called(filters, associations, pagination, ordering)

	var r0 models.TransportationServiceProviderPerformances
	if rf, ok := ret.Get(0).(func([]services.QueryFilter, services.QueryAssociations, services.Pagination, services.QueryOrder) models.TransportationServiceProviderPerformances); ok {
		r0 = rf(filters, associations, pagination, ordering)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.TransportationServiceProviderPerformances)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]services.QueryFilter, services.QueryAssociations, services.Pagination, services.QueryOrder) error); ok {
		r1 = rf(filters, associations, pagination, ordering)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
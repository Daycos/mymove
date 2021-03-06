// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	validate "github.com/gobuffalo/validate"
)

// PaymentRequestStatusUpdater is an autogenerated mock type for the PaymentRequestStatusUpdater type
type PaymentRequestStatusUpdater struct {
	mock.Mock
}

// UpdatePaymentRequestStatus provides a mock function with given fields: _a0
func (_m *PaymentRequestStatusUpdater) UpdatePaymentRequestStatus(_a0 *models.PaymentRequest) (*validate.Errors, error) {
	ret := _m.Called(_a0)

	var r0 *validate.Errors
	if rf, ok := ret.Get(0).(func(*models.PaymentRequest) *validate.Errors); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*validate.Errors)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.PaymentRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

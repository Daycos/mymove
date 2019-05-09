// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/transcom/mymove/pkg/models"

import uuid "github.com/gofrs/uuid"

// AccessCodeClaimer is an autogenerated mock type for the AccessCodeClaimer type
type AccessCodeClaimer struct {
	mock.Mock
}

// ClaimAccessCode provides a mock function with given fields: code, serviceMemberID
func (_m *AccessCodeClaimer) ClaimAccessCode(code string, serviceMemberID uuid.UUID) (*models.AccessCode, error) {
	ret := _m.Called(code, serviceMemberID)

	var r0 *models.AccessCode
	if rf, ok := ret.Get(0).(func(string, uuid.UUID) *models.AccessCode); ok {
		r0 = rf(code, serviceMemberID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AccessCode)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uuid.UUID) error); ok {
		r1 = rf(code, serviceMemberID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

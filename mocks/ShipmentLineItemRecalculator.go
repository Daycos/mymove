// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import auth "github.com/transcom/mymove/pkg/auth"
import mock "github.com/stretchr/testify/mock"
import models "github.com/transcom/mymove/pkg/models"
import route "github.com/transcom/mymove/pkg/route"

import uuid "github.com/gofrs/uuid"

// ShipmentLineItemRecalculator is an autogenerated mock type for the ShipmentLineItemRecalculator type
type ShipmentLineItemRecalculator struct {
	mock.Mock
}

// RecalculateShipmentLineItems provides a mock function with given fields: shipmentID, session, _a2
func (_m *ShipmentLineItemRecalculator) RecalculateShipmentLineItems(shipmentID uuid.UUID, session *auth.Session, _a2 route.Planner) ([]models.ShipmentLineItem, error) {
	ret := _m.Called(shipmentID, session, _a2)

	var r0 []models.ShipmentLineItem
	if rf, ok := ret.Get(0).(func(uuid.UUID, *auth.Session, route.Planner) []models.ShipmentLineItem); ok {
		r0 = rf(shipmentID, session, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ShipmentLineItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, *auth.Session, route.Planner) error); ok {
		r1 = rf(shipmentID, session, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
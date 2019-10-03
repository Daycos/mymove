// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	auth "github.com/transcom/mymove/pkg/auth"
	internalmessages "github.com/transcom/mymove/pkg/gen/internalmessages"

	models "github.com/transcom/mymove/pkg/models"

	validate "github.com/gobuffalo/validate"
)

// Updater is an autogenerated mock type for the Updater type
type Updater struct {
	mock.Mock
}

// Update provides a mock function with given fields: moveDocumentPayload, moveDoc, session
func (_m *Updater) Update(moveDocumentPayload *internalmessages.MoveDocumentPayload, moveDoc *models.MoveDocument, session *auth.Session) (*models.MoveDocument, *validate.Errors, error) {
	ret := _m.Called(moveDocumentPayload, moveDoc, session)

	var r0 *models.MoveDocument
	if rf, ok := ret.Get(0).(func(*internalmessages.MoveDocumentPayload, *models.MoveDocument, *auth.Session) *models.MoveDocument); ok {
		r0 = rf(moveDocumentPayload, moveDoc, session)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MoveDocument)
		}
	}

	var r1 *validate.Errors
	if rf, ok := ret.Get(1).(func(*internalmessages.MoveDocumentPayload, *models.MoveDocument, *auth.Session) *validate.Errors); ok {
		r1 = rf(moveDocumentPayload, moveDoc, session)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*validate.Errors)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*internalmessages.MoveDocumentPayload, *models.MoveDocument, *auth.Session) error); ok {
		r2 = rf(moveDocumentPayload, moveDoc, session)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

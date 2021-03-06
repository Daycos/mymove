// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MoveTaskOrder move task order
// swagger:model MoveTaskOrder
type MoveTaskOrder struct {

	// created at
	// Format: date
	CreatedAt strfmt.Date `json:"createdAt,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// is available to prime
	IsAvailableToPrime *bool `json:"isAvailableToPrime,omitempty"`

	// is canceled
	IsCanceled *bool `json:"isCanceled,omitempty"`

	// move order
	MoveOrder *MoveOrder `json:"moveOrder,omitempty"`

	// move order ID
	// Format: uuid
	MoveOrderID strfmt.UUID `json:"moveOrderID,omitempty"`

	// mto service items
	// Required: true
	MtoServiceItems MTOServiceItems `json:"mto_service_items"`

	// mto shipments
	// Required: true
	MtoShipments MTOShipments `json:"mto_shipments"`

	// payment requests
	// Required: true
	PaymentRequests PaymentRequests `json:"payment_requests"`

	// ppm estimated weight
	PpmEstimatedWeight int64 `json:"ppm_estimated_weight,omitempty"`

	// ppm type
	// Enum: [FULL PARTIAL]
	PpmType string `json:"ppm_type,omitempty"`

	// reference Id
	ReferenceID string `json:"referenceId,omitempty"`

	// updated at
	// Format: date
	UpdatedAt strfmt.Date `json:"updatedAt,omitempty"`
}

// Validate validates this move task order
func (m *MoveTaskOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtoServiceItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtoShipments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePpmType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveTaskOrder) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateMoveOrder(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveOrder) { // not required
		return nil
	}

	if m.MoveOrder != nil {
		if err := m.MoveOrder.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moveOrder")
			}
			return err
		}
	}

	return nil
}

func (m *MoveTaskOrder) validateMoveOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveOrderID) { // not required
		return nil
	}

	if err := validate.FormatOf("moveOrderID", "body", "uuid", m.MoveOrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateMtoServiceItems(formats strfmt.Registry) error {

	if err := validate.Required("mto_service_items", "body", m.MtoServiceItems); err != nil {
		return err
	}

	if err := m.MtoServiceItems.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mto_service_items")
		}
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateMtoShipments(formats strfmt.Registry) error {

	if err := validate.Required("mto_shipments", "body", m.MtoShipments); err != nil {
		return err
	}

	if err := m.MtoShipments.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mto_shipments")
		}
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validatePaymentRequests(formats strfmt.Registry) error {

	if err := validate.Required("payment_requests", "body", m.PaymentRequests); err != nil {
		return err
	}

	if err := m.PaymentRequests.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("payment_requests")
		}
		return err
	}

	return nil
}

var moveTaskOrderTypePpmTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FULL","PARTIAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		moveTaskOrderTypePpmTypePropEnum = append(moveTaskOrderTypePpmTypePropEnum, v)
	}
}

const (

	// MoveTaskOrderPpmTypeFULL captures enum value "FULL"
	MoveTaskOrderPpmTypeFULL string = "FULL"

	// MoveTaskOrderPpmTypePARTIAL captures enum value "PARTIAL"
	MoveTaskOrderPpmTypePARTIAL string = "PARTIAL"
)

// prop value enum
func (m *MoveTaskOrder) validatePpmTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, moveTaskOrderTypePpmTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MoveTaskOrder) validatePpmType(formats strfmt.Registry) error {

	if swag.IsZero(m.PpmType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePpmTypeEnum("ppm_type", "body", m.PpmType); err != nil {
		return err
	}

	return nil
}

func (m *MoveTaskOrder) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoveTaskOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveTaskOrder) UnmarshalBinary(b []byte) error {
	var res MoveTaskOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

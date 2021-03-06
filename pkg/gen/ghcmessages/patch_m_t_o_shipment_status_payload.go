// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchMTOShipmentStatusPayload patch m t o shipment status payload
// swagger:model PatchMTOShipmentStatusPayload
type PatchMTOShipmentStatusPayload struct {

	// rejection reason
	RejectionReason *string `json:"rejectionReason,omitempty"`

	// status
	// Enum: [REJECTED APPROVED SUBMITTED]
	Status string `json:"status,omitempty"`
}

// Validate validates this patch m t o shipment status payload
func (m *PatchMTOShipmentStatusPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var patchMTOShipmentStatusPayloadTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REJECTED","APPROVED","SUBMITTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchMTOShipmentStatusPayloadTypeStatusPropEnum = append(patchMTOShipmentStatusPayloadTypeStatusPropEnum, v)
	}
}

const (

	// PatchMTOShipmentStatusPayloadStatusREJECTED captures enum value "REJECTED"
	PatchMTOShipmentStatusPayloadStatusREJECTED string = "REJECTED"

	// PatchMTOShipmentStatusPayloadStatusAPPROVED captures enum value "APPROVED"
	PatchMTOShipmentStatusPayloadStatusAPPROVED string = "APPROVED"

	// PatchMTOShipmentStatusPayloadStatusSUBMITTED captures enum value "SUBMITTED"
	PatchMTOShipmentStatusPayloadStatusSUBMITTED string = "SUBMITTED"
)

// prop value enum
func (m *PatchMTOShipmentStatusPayload) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, patchMTOShipmentStatusPayloadTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PatchMTOShipmentStatusPayload) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchMTOShipmentStatusPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchMTOShipmentStatusPayload) UnmarshalBinary(b []byte) error {
	var res PatchMTOShipmentStatusPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// MethodOfReceipt Method of Receipt
// swagger:model MethodOfReceipt
type MethodOfReceipt string

const (

	// MethodOfReceiptMILPAY captures enum value "MIL_PAY"
	MethodOfReceiptMILPAY MethodOfReceipt = "MIL_PAY"

	// MethodOfReceiptOTHERDD captures enum value "OTHER_DD"
	MethodOfReceiptOTHERDD MethodOfReceipt = "OTHER_DD"

	// MethodOfReceiptGTCC captures enum value "GTCC"
	MethodOfReceiptGTCC MethodOfReceipt = "GTCC"
)

// for schema
var methodOfReceiptEnum []interface{}

func init() {
	var res []MethodOfReceipt
	if err := json.Unmarshal([]byte(`["MIL_PAY","OTHER_DD","GTCC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		methodOfReceiptEnum = append(methodOfReceiptEnum, v)
	}
}

func (m MethodOfReceipt) validateMethodOfReceiptEnum(path, location string, value MethodOfReceipt) error {
	if err := validate.Enum(path, location, value, methodOfReceiptEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this method of receipt
func (m MethodOfReceipt) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMethodOfReceiptEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

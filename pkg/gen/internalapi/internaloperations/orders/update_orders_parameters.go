// Code generated by go-swagger; DO NOT EDIT.

package orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	internalmessages "github.com/transcom/mymove/pkg/gen/internalmessages"
)

// NewUpdateOrdersParams creates a new UpdateOrdersParams object
// no default values defined in spec.
func NewUpdateOrdersParams() UpdateOrdersParams {

	return UpdateOrdersParams{}
}

// UpdateOrdersParams contains all the bound params for the update orders operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateOrders
type UpdateOrdersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*UUID of the orders model
	  Required: true
	  In: path
	*/
	OrdersID strfmt.UUID
	/*
	  Required: true
	  In: body
	*/
	UpdateOrders *internalmessages.CreateUpdateOrders
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateOrdersParams() beforehand.
func (o *UpdateOrdersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rOrdersID, rhkOrdersID, _ := route.Params.GetOK("ordersId")
	if err := o.bindOrdersID(rOrdersID, rhkOrdersID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body internalmessages.CreateUpdateOrders
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("updateOrders", "body"))
			} else {
				res = append(res, errors.NewParseError("updateOrders", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.UpdateOrders = &body
			}
		}
	} else {
		res = append(res, errors.Required("updateOrders", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindOrdersID binds and validates parameter OrdersID from path.
func (o *UpdateOrdersParams) bindOrdersID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("ordersId", "path", "strfmt.UUID", raw)
	}
	o.OrdersID = *(value.(*strfmt.UUID))

	if err := o.validateOrdersID(formats); err != nil {
		return err
	}

	return nil
}

// validateOrdersID carries on validations for parameter OrdersID
func (o *UpdateOrdersParams) validateOrdersID(formats strfmt.Registry) error {

	if err := validate.FormatOf("ordersId", "path", "uuid", o.OrdersID.String(), formats); err != nil {
		return err
	}
	return nil
}

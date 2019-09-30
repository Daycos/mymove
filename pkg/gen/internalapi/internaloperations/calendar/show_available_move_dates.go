// Code generated by go-swagger; DO NOT EDIT.

package calendar

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ShowAvailableMoveDatesHandlerFunc turns a function with the right signature into a show available move dates handler
type ShowAvailableMoveDatesHandlerFunc func(ShowAvailableMoveDatesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ShowAvailableMoveDatesHandlerFunc) Handle(params ShowAvailableMoveDatesParams) middleware.Responder {
	return fn(params)
}

// ShowAvailableMoveDatesHandler interface for that can handle valid show available move dates params
type ShowAvailableMoveDatesHandler interface {
	Handle(ShowAvailableMoveDatesParams) middleware.Responder
}

// NewShowAvailableMoveDates creates a new http.Handler for the show available move dates operation
func NewShowAvailableMoveDates(ctx *middleware.Context, handler ShowAvailableMoveDatesHandler) *ShowAvailableMoveDates {
	return &ShowAvailableMoveDates{Context: ctx, Handler: handler}
}

/*ShowAvailableMoveDates swagger:route GET /calendar/available_move_dates calendar showAvailableMoveDates

Returns available dates for the move calendar

Returns available dates for the move calendar

*/
type ShowAvailableMoveDates struct {
	Context *middleware.Context
	Handler ShowAvailableMoveDatesHandler
}

func (o *ShowAvailableMoveDates) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewShowAvailableMoveDatesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

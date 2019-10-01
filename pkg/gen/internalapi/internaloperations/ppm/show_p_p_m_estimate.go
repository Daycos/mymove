// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ShowPPMEstimateHandlerFunc turns a function with the right signature into a show p p m estimate handler
type ShowPPMEstimateHandlerFunc func(ShowPPMEstimateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ShowPPMEstimateHandlerFunc) Handle(params ShowPPMEstimateParams) middleware.Responder {
	return fn(params)
}

// ShowPPMEstimateHandler interface for that can handle valid show p p m estimate params
type ShowPPMEstimateHandler interface {
	Handle(ShowPPMEstimateParams) middleware.Responder
}

// NewShowPPMEstimate creates a new http.Handler for the show p p m estimate operation
func NewShowPPMEstimate(ctx *middleware.Context, handler ShowPPMEstimateHandler) *ShowPPMEstimate {
	return &ShowPPMEstimate{Context: ctx, Handler: handler}
}

/*ShowPPMEstimate swagger:route GET /estimates/ppm ppm showPPMEstimate

Return a PPM cost estimate

Calculates a reimbursement range for a PPM move (excluding SIT)

*/
type ShowPPMEstimate struct {
	Context *middleware.Context
	Handler ShowPPMEstimateHandler
}

func (o *ShowPPMEstimate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewShowPPMEstimateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
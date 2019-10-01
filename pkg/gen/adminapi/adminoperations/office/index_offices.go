// Code generated by go-swagger; DO NOT EDIT.

package office

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// IndexOfficesHandlerFunc turns a function with the right signature into a index offices handler
type IndexOfficesHandlerFunc func(IndexOfficesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn IndexOfficesHandlerFunc) Handle(params IndexOfficesParams) middleware.Responder {
	return fn(params)
}

// IndexOfficesHandler interface for that can handle valid index offices params
type IndexOfficesHandler interface {
	Handle(IndexOfficesParams) middleware.Responder
}

// NewIndexOffices creates a new http.Handler for the index offices operation
func NewIndexOffices(ctx *middleware.Context, handler IndexOfficesHandler) *IndexOffices {
	return &IndexOffices{Context: ctx, Handler: handler}
}

/*IndexOffices swagger:route GET /offices office indexOffices

List transportation offices

Returns a list of transportation offices

*/
type IndexOffices struct {
	Context *middleware.Context
	Handler IndexOfficesHandler
}

func (o *IndexOffices) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewIndexOfficesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
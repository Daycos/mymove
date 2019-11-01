// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateMoveTaskOrderStatusHandlerFunc turns a function with the right signature into a update move task order status handler
type UpdateMoveTaskOrderStatusHandlerFunc func(UpdateMoveTaskOrderStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateMoveTaskOrderStatusHandlerFunc) Handle(params UpdateMoveTaskOrderStatusParams) middleware.Responder {
	return fn(params)
}

// UpdateMoveTaskOrderStatusHandler interface for that can handle valid update move task order status params
type UpdateMoveTaskOrderStatusHandler interface {
	Handle(UpdateMoveTaskOrderStatusParams) middleware.Responder
}

// NewUpdateMoveTaskOrderStatus creates a new http.Handler for the update move task order status operation
func NewUpdateMoveTaskOrderStatus(ctx *middleware.Context, handler UpdateMoveTaskOrderStatusHandler) *UpdateMoveTaskOrderStatus {
	return &UpdateMoveTaskOrderStatus{Context: ctx, Handler: handler}
}

/*UpdateMoveTaskOrderStatus swagger:route PATCH /move-task-orders/{moveTaskOrderID}/status moveTaskOrder updateMoveTaskOrderStatus

Change the status of a move order

Changes move order status

*/
type UpdateMoveTaskOrderStatus struct {
	Context *middleware.Context
	Handler UpdateMoveTaskOrderStatusHandler
}

func (o *UpdateMoveTaskOrderStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateMoveTaskOrderStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
// Code generated by go-swagger; DO NOT EDIT.

package service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteServiceItemHandlerFunc turns a function with the right signature into a delete service item handler
type DeleteServiceItemHandlerFunc func(DeleteServiceItemParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteServiceItemHandlerFunc) Handle(params DeleteServiceItemParams) middleware.Responder {
	return fn(params)
}

// DeleteServiceItemHandler interface for that can handle valid delete service item params
type DeleteServiceItemHandler interface {
	Handle(DeleteServiceItemParams) middleware.Responder
}

// NewDeleteServiceItem creates a new http.Handler for the delete service item operation
func NewDeleteServiceItem(ctx *middleware.Context, handler DeleteServiceItemHandler) *DeleteServiceItem {
	return &DeleteServiceItem{Context: ctx, Handler: handler}
}

/*DeleteServiceItem swagger:route DELETE /move-task-orders/{moveTaskOrderID}/service-items/{serviceItemID} serviceItem deleteServiceItem

Deletes a line item by ID for a move order by ID

Deletes a line item by ID for a move order by ID

*/
type DeleteServiceItem struct {
	Context *middleware.Context
	Handler DeleteServiceItemHandler
}

func (o *DeleteServiceItem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteServiceItemParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
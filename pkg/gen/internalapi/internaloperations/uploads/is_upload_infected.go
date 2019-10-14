// Code generated by go-swagger; DO NOT EDIT.

package uploads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// IsUploadInfectedHandlerFunc turns a function with the right signature into a is upload infected handler
type IsUploadInfectedHandlerFunc func(IsUploadInfectedParams) middleware.Responder

// Handle executing the request and returning a response
func (fn IsUploadInfectedHandlerFunc) Handle(params IsUploadInfectedParams) middleware.Responder {
	return fn(params)
}

// IsUploadInfectedHandler interface for that can handle valid is upload infected params
type IsUploadInfectedHandler interface {
	Handle(IsUploadInfectedParams) middleware.Responder
}

// NewIsUploadInfected creates a new http.Handler for the is upload infected operation
func NewIsUploadInfected(ctx *middleware.Context, handler IsUploadInfectedHandler) *IsUploadInfected {
	return &IsUploadInfected{Context: ctx, Handler: handler}
}

/*IsUploadInfected swagger:route GET /uploads/{uploadId}/is_infected uploads isUploadInfected

Returns boolean as to whether the upload is infected

Returns boolean as to whether the upload is infected

*/
type IsUploadInfected struct {
	Context *middleware.Context
	Handler IsUploadInfectedHandler
}

func (o *IsUploadInfected) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewIsUploadInfectedParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

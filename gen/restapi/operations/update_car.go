// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateCarHandlerFunc turns a function with the right signature into a update car handler
type UpdateCarHandlerFunc func(UpdateCarParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateCarHandlerFunc) Handle(params UpdateCarParams) middleware.Responder {
	return fn(params)
}

// UpdateCarHandler interface for that can handle valid update car params
type UpdateCarHandler interface {
	Handle(UpdateCarParams) middleware.Responder
}

// NewUpdateCar creates a new http.Handler for the update car operation
func NewUpdateCar(ctx *middleware.Context, handler UpdateCarHandler) *UpdateCar {
	return &UpdateCar{Context: ctx, Handler: handler}
}

/*
	UpdateCar swagger:route PUT /car updateCar

UpdateCar update car API
*/
type UpdateCar struct {
	Context *middleware.Context
	Handler UpdateCarHandler
}

func (o *UpdateCar) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateCarParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

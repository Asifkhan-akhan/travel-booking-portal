// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteBookingHandlerFunc turns a function with the right signature into a delete booking handler
type DeleteBookingHandlerFunc func(DeleteBookingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteBookingHandlerFunc) Handle(params DeleteBookingParams) middleware.Responder {
	return fn(params)
}

// DeleteBookingHandler interface for that can handle valid delete booking params
type DeleteBookingHandler interface {
	Handle(DeleteBookingParams) middleware.Responder
}

// NewDeleteBooking creates a new http.Handler for the delete booking operation
func NewDeleteBooking(ctx *middleware.Context, handler DeleteBookingHandler) *DeleteBooking {
	return &DeleteBooking{Context: ctx, Handler: handler}
}

/*
	DeleteBooking swagger:route DELETE /booking/{bookingID} deleteBooking

DeleteBooking delete booking API
*/
type DeleteBooking struct {
	Context *middleware.Context
	Handler DeleteBookingHandler
}

func (o *DeleteBooking) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteBookingParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

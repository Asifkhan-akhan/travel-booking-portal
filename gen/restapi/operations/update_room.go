// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateRoomHandlerFunc turns a function with the right signature into a update room handler
type UpdateRoomHandlerFunc func(UpdateRoomParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateRoomHandlerFunc) Handle(params UpdateRoomParams) middleware.Responder {
	return fn(params)
}

// UpdateRoomHandler interface for that can handle valid update room params
type UpdateRoomHandler interface {
	Handle(UpdateRoomParams) middleware.Responder
}

// NewUpdateRoom creates a new http.Handler for the update room operation
func NewUpdateRoom(ctx *middleware.Context, handler UpdateRoomHandler) *UpdateRoom {
	return &UpdateRoom{Context: ctx, Handler: handler}
}

/*
	UpdateRoom swagger:route PUT /room updateRoom

UpdateRoom update room API
*/
type UpdateRoom struct {
	Context *middleware.Context
	Handler UpdateRoomHandler
}

func (o *UpdateRoom) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateRoomParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListRoomsHandlerFunc turns a function with the right signature into a list rooms handler
type ListRoomsHandlerFunc func(ListRoomsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListRoomsHandlerFunc) Handle(params ListRoomsParams) middleware.Responder {
	return fn(params)
}

// ListRoomsHandler interface for that can handle valid list rooms params
type ListRoomsHandler interface {
	Handle(ListRoomsParams) middleware.Responder
}

// NewListRooms creates a new http.Handler for the list rooms operation
func NewListRooms(ctx *middleware.Context, handler ListRoomsHandler) *ListRooms {
	return &ListRooms{Context: ctx, Handler: handler}
}

/*
	ListRooms swagger:route GET /room listRooms

ListRooms list rooms API
*/
type ListRooms struct {
	Context *middleware.Context
	Handler ListRoomsHandler
}

func (o *ListRooms) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListRoomsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

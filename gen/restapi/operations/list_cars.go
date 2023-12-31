// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListCarsHandlerFunc turns a function with the right signature into a list cars handler
type ListCarsHandlerFunc func(ListCarsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListCarsHandlerFunc) Handle(params ListCarsParams) middleware.Responder {
	return fn(params)
}

// ListCarsHandler interface for that can handle valid list cars params
type ListCarsHandler interface {
	Handle(ListCarsParams) middleware.Responder
}

// NewListCars creates a new http.Handler for the list cars operation
func NewListCars(ctx *middleware.Context, handler ListCarsHandler) *ListCars {
	return &ListCars{Context: ctx, Handler: handler}
}

/*
	ListCars swagger:route GET /car listCars

ListCars list cars API
*/
type ListCars struct {
	Context *middleware.Context
	Handler ListCarsHandler
}

func (o *ListCars) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListCarsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

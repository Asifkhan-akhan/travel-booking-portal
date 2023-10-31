// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteCarParams creates a new DeleteCarParams object
//
// There are no default values defined in the spec.
func NewDeleteCarParams() DeleteCarParams {

	return DeleteCarParams{}
}

// DeleteCarParams contains all the bound params for the delete car operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteCar
type DeleteCarParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of the Car to delete
	  Required: true
	  In: path
	*/
	CarID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteCarParams() beforehand.
func (o *DeleteCarParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rCarID, rhkCarID, _ := route.Params.GetOK("carID")
	if err := o.bindCarID(rCarID, rhkCarID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCarID binds and validates parameter CarID from path.
func (o *DeleteCarParams) bindCarID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("carID", "path", "int64", raw)
	}
	o.CarID = value

	return nil
}

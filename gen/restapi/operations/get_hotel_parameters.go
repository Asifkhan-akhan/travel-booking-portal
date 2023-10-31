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

// NewGetHotelParams creates a new GetHotelParams object
//
// There are no default values defined in the spec.
func NewGetHotelParams() GetHotelParams {

	return GetHotelParams{}
}

// GetHotelParams contains all the bound params for the get hotel operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetHotel
type GetHotelParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of the Hotel to retrieve
	  Required: true
	  In: path
	*/
	HotelID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetHotelParams() beforehand.
func (o *GetHotelParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rHotelID, rhkHotelID, _ := route.Params.GetOK("hotelID")
	if err := o.bindHotelID(rHotelID, rhkHotelID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindHotelID binds and validates parameter HotelID from path.
func (o *GetHotelParams) bindHotelID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("hotelID", "path", "int64", raw)
	}
	o.HotelID = value

	return nil
}

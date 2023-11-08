// Code generated by go-swagger; DO NOT EDIT.

package bookings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewListBookingsByDateTimeParams creates a new ListBookingsByDateTimeParams object
//
// There are no default values defined in the spec.
func NewListBookingsByDateTimeParams() ListBookingsByDateTimeParams {

	return ListBookingsByDateTimeParams{}
}

// ListBookingsByDateTimeParams contains all the bound params for the list bookings by date time operation
// typically these are obtained from a http.Request
//
// swagger:parameters ListBookingsByDateTime
type ListBookingsByDateTimeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Date and time for filtering bookings
	  Required: true
	  In: path
	*/
	DateTime string
	/*User ID for filtering bookings
	  Required: true
	  In: path
	*/
	UserID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListBookingsByDateTimeParams() beforehand.
func (o *ListBookingsByDateTimeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rDateTime, rhkDateTime, _ := route.Params.GetOK("dateTime")
	if err := o.bindDateTime(rDateTime, rhkDateTime, route.Formats); err != nil {
		res = append(res, err)
	}

	rUserID, rhkUserID, _ := route.Params.GetOK("userID")
	if err := o.bindUserID(rUserID, rhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDateTime binds and validates parameter DateTime from path.
func (o *ListBookingsByDateTimeParams) bindDateTime(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.DateTime = raw

	return nil
}

// bindUserID binds and validates parameter UserID from path.
func (o *ListBookingsByDateTimeParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("userID", "path", "int64", raw)
	}
	o.UserID = value

	return nil
}

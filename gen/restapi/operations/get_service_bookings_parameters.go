// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetServiceBookingsParams creates a new GetServiceBookingsParams object
//
// There are no default values defined in the spec.
func NewGetServiceBookingsParams() GetServiceBookingsParams {

	return GetServiceBookingsParams{}
}

// GetServiceBookingsParams contains all the bound params for the get service bookings operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetServiceBookings
type GetServiceBookingsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Start date for filtering bookings
	  In: query
	*/
	FromDate *string
	/*ID of the service to retrieve bookings
	  Required: true
	  In: path
	*/
	ServiceID int64
	/*End date for filtering bookings
	  In: query
	*/
	ToDate *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetServiceBookingsParams() beforehand.
func (o *GetServiceBookingsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFromDate, qhkFromDate, _ := qs.GetOK("FromDate")
	if err := o.bindFromDate(qFromDate, qhkFromDate, route.Formats); err != nil {
		res = append(res, err)
	}

	rServiceID, rhkServiceID, _ := route.Params.GetOK("ServiceID")
	if err := o.bindServiceID(rServiceID, rhkServiceID, route.Formats); err != nil {
		res = append(res, err)
	}

	qToDate, qhkToDate, _ := qs.GetOK("ToDate")
	if err := o.bindToDate(qToDate, qhkToDate, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFromDate binds and validates parameter FromDate from query.
func (o *GetServiceBookingsParams) bindFromDate(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.FromDate = &raw

	return nil
}

// bindServiceID binds and validates parameter ServiceID from path.
func (o *GetServiceBookingsParams) bindServiceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("ServiceID", "path", "int64", raw)
	}
	o.ServiceID = value

	return nil
}

// bindToDate binds and validates parameter ToDate from query.
func (o *GetServiceBookingsParams) bindToDate(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.ToDate = &raw

	return nil
}

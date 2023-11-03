// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"travel-booking-portal/gen/models"
)

// NewCreateCarParams creates a new CreateCarParams object
//
// There are no default values defined in the spec.
func NewCreateCarParams() CreateCarParams {

	return CreateCarParams{}
}

// CreateCarParams contains all the bound params for the create car operation
// typically these are obtained from a http.Request
//
// swagger:parameters CreateCar
type CreateCarParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*CarRental object to create
	  Required: true
	  In: body
	*/
	Car *models.CarRental
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateCarParams() beforehand.
func (o *CreateCarParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.CarRental
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("car", "body", ""))
			} else {
				res = append(res, errors.NewParseError("car", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Car = &body
			}
		}
	} else {
		res = append(res, errors.Required("car", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
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

// NewCreateOrUpdateRoomParams creates a new CreateOrUpdateRoomParams object
//
// There are no default values defined in the spec.
func NewCreateOrUpdateRoomParams() CreateOrUpdateRoomParams {

	return CreateOrUpdateRoomParams{}
}

// CreateOrUpdateRoomParams contains all the bound params for the create or update room operation
// typically these are obtained from a http.Request
//
// swagger:parameters CreateOrUpdateRoom
type CreateOrUpdateRoomParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Room object to create or update
	  Required: true
	  In: body
	*/
	Room *models.Room
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateOrUpdateRoomParams() beforehand.
func (o *CreateOrUpdateRoomParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Room
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("room", "body", ""))
			} else {
				res = append(res, errors.NewParseError("room", "body", "", err))
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
				o.Room = &body
			}
		}
	} else {
		res = append(res, errors.Required("room", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

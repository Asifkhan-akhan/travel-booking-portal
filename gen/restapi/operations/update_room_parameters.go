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

// NewUpdateRoomParams creates a new UpdateRoomParams object
//
// There are no default values defined in the spec.
func NewUpdateRoomParams() UpdateRoomParams {

	return UpdateRoomParams{}
}

// UpdateRoomParams contains all the bound params for the update room operation
// typically these are obtained from a http.Request
//
// swagger:parameters UpdateRoom
type UpdateRoomParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Room object to  update
	  Required: true
	  In: body
	*/
	Room *models.Room
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateRoomParams() beforehand.
func (o *UpdateRoomParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

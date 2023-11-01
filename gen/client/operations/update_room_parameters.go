// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"travel-booking-portal/gen/models"
)

// NewUpdateRoomParams creates a new UpdateRoomParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRoomParams() *UpdateRoomParams {
	return &UpdateRoomParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRoomParamsWithTimeout creates a new UpdateRoomParams object
// with the ability to set a timeout on a request.
func NewUpdateRoomParamsWithTimeout(timeout time.Duration) *UpdateRoomParams {
	return &UpdateRoomParams{
		timeout: timeout,
	}
}

// NewUpdateRoomParamsWithContext creates a new UpdateRoomParams object
// with the ability to set a context for a request.
func NewUpdateRoomParamsWithContext(ctx context.Context) *UpdateRoomParams {
	return &UpdateRoomParams{
		Context: ctx,
	}
}

// NewUpdateRoomParamsWithHTTPClient creates a new UpdateRoomParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRoomParamsWithHTTPClient(client *http.Client) *UpdateRoomParams {
	return &UpdateRoomParams{
		HTTPClient: client,
	}
}

/*
UpdateRoomParams contains all the parameters to send to the API endpoint

	for the update room operation.

	Typically these are written to a http.Request.
*/
type UpdateRoomParams struct {

	/* Room.

	   Room object to  update
	*/
	Room *models.Room

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update room params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRoomParams) WithDefaults() *UpdateRoomParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update room params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRoomParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update room params
func (o *UpdateRoomParams) WithTimeout(timeout time.Duration) *UpdateRoomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update room params
func (o *UpdateRoomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update room params
func (o *UpdateRoomParams) WithContext(ctx context.Context) *UpdateRoomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update room params
func (o *UpdateRoomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update room params
func (o *UpdateRoomParams) WithHTTPClient(client *http.Client) *UpdateRoomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update room params
func (o *UpdateRoomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoom adds the room to the update room params
func (o *UpdateRoomParams) WithRoom(room *models.Room) *UpdateRoomParams {
	o.SetRoom(room)
	return o
}

// SetRoom adds the room to the update room params
func (o *UpdateRoomParams) SetRoom(room *models.Room) {
	o.Room = room
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRoomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Room != nil {
		if err := r.SetBodyParam(o.Room); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

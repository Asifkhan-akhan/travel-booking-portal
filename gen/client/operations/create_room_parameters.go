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

// NewCreateRoomParams creates a new CreateRoomParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRoomParams() *CreateRoomParams {
	return &CreateRoomParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRoomParamsWithTimeout creates a new CreateRoomParams object
// with the ability to set a timeout on a request.
func NewCreateRoomParamsWithTimeout(timeout time.Duration) *CreateRoomParams {
	return &CreateRoomParams{
		timeout: timeout,
	}
}

// NewCreateRoomParamsWithContext creates a new CreateRoomParams object
// with the ability to set a context for a request.
func NewCreateRoomParamsWithContext(ctx context.Context) *CreateRoomParams {
	return &CreateRoomParams{
		Context: ctx,
	}
}

// NewCreateRoomParamsWithHTTPClient creates a new CreateRoomParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRoomParamsWithHTTPClient(client *http.Client) *CreateRoomParams {
	return &CreateRoomParams{
		HTTPClient: client,
	}
}

/*
CreateRoomParams contains all the parameters to send to the API endpoint

	for the create room operation.

	Typically these are written to a http.Request.
*/
type CreateRoomParams struct {

	/* Room.

	   Room object to create
	*/
	Room *models.Room

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create room params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRoomParams) WithDefaults() *CreateRoomParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create room params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRoomParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create room params
func (o *CreateRoomParams) WithTimeout(timeout time.Duration) *CreateRoomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create room params
func (o *CreateRoomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create room params
func (o *CreateRoomParams) WithContext(ctx context.Context) *CreateRoomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create room params
func (o *CreateRoomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create room params
func (o *CreateRoomParams) WithHTTPClient(client *http.Client) *CreateRoomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create room params
func (o *CreateRoomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoom adds the room to the create room params
func (o *CreateRoomParams) WithRoom(room *models.Room) *CreateRoomParams {
	o.SetRoom(room)
	return o
}

// SetRoom adds the room to the create room params
func (o *CreateRoomParams) SetRoom(room *models.Room) {
	o.Room = room
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRoomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
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
)

// NewGetRoomsParams creates a new GetRoomsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRoomsParams() *GetRoomsParams {
	return &GetRoomsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoomsParamsWithTimeout creates a new GetRoomsParams object
// with the ability to set a timeout on a request.
func NewGetRoomsParamsWithTimeout(timeout time.Duration) *GetRoomsParams {
	return &GetRoomsParams{
		timeout: timeout,
	}
}

// NewGetRoomsParamsWithContext creates a new GetRoomsParams object
// with the ability to set a context for a request.
func NewGetRoomsParamsWithContext(ctx context.Context) *GetRoomsParams {
	return &GetRoomsParams{
		Context: ctx,
	}
}

// NewGetRoomsParamsWithHTTPClient creates a new GetRoomsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRoomsParamsWithHTTPClient(client *http.Client) *GetRoomsParams {
	return &GetRoomsParams{
		HTTPClient: client,
	}
}

/*
GetRoomsParams contains all the parameters to send to the API endpoint

	for the get rooms operation.

	Typically these are written to a http.Request.
*/
type GetRoomsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get rooms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoomsParams) WithDefaults() *GetRoomsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get rooms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoomsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get rooms params
func (o *GetRoomsParams) WithTimeout(timeout time.Duration) *GetRoomsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rooms params
func (o *GetRoomsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rooms params
func (o *GetRoomsParams) WithContext(ctx context.Context) *GetRoomsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rooms params
func (o *GetRoomsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rooms params
func (o *GetRoomsParams) WithHTTPClient(client *http.Client) *GetRoomsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rooms params
func (o *GetRoomsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoomsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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
	"github.com/go-openapi/swag"
)

// NewDeleteHotelParams creates a new DeleteHotelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteHotelParams() *DeleteHotelParams {
	return &DeleteHotelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteHotelParamsWithTimeout creates a new DeleteHotelParams object
// with the ability to set a timeout on a request.
func NewDeleteHotelParamsWithTimeout(timeout time.Duration) *DeleteHotelParams {
	return &DeleteHotelParams{
		timeout: timeout,
	}
}

// NewDeleteHotelParamsWithContext creates a new DeleteHotelParams object
// with the ability to set a context for a request.
func NewDeleteHotelParamsWithContext(ctx context.Context) *DeleteHotelParams {
	return &DeleteHotelParams{
		Context: ctx,
	}
}

// NewDeleteHotelParamsWithHTTPClient creates a new DeleteHotelParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteHotelParamsWithHTTPClient(client *http.Client) *DeleteHotelParams {
	return &DeleteHotelParams{
		HTTPClient: client,
	}
}

/*
DeleteHotelParams contains all the parameters to send to the API endpoint

	for the delete hotel operation.

	Typically these are written to a http.Request.
*/
type DeleteHotelParams struct {

	/* HotelID.

	   ID of the Hotel to delete
	*/
	HotelID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete hotel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteHotelParams) WithDefaults() *DeleteHotelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete hotel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteHotelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete hotel params
func (o *DeleteHotelParams) WithTimeout(timeout time.Duration) *DeleteHotelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete hotel params
func (o *DeleteHotelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete hotel params
func (o *DeleteHotelParams) WithContext(ctx context.Context) *DeleteHotelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete hotel params
func (o *DeleteHotelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete hotel params
func (o *DeleteHotelParams) WithHTTPClient(client *http.Client) *DeleteHotelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete hotel params
func (o *DeleteHotelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHotelID adds the hotelID to the delete hotel params
func (o *DeleteHotelParams) WithHotelID(hotelID int64) *DeleteHotelParams {
	o.SetHotelID(hotelID)
	return o
}

// SetHotelID adds the hotelId to the delete hotel params
func (o *DeleteHotelParams) SetHotelID(hotelID int64) {
	o.HotelID = hotelID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteHotelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param hotelID
	if err := r.SetPathParam("hotelID", swag.FormatInt64(o.HotelID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

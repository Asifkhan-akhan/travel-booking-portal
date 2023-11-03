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

// NewDeleteBookingParams creates a new DeleteBookingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBookingParams() *DeleteBookingParams {
	return &DeleteBookingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBookingParamsWithTimeout creates a new DeleteBookingParams object
// with the ability to set a timeout on a request.
func NewDeleteBookingParamsWithTimeout(timeout time.Duration) *DeleteBookingParams {
	return &DeleteBookingParams{
		timeout: timeout,
	}
}

// NewDeleteBookingParamsWithContext creates a new DeleteBookingParams object
// with the ability to set a context for a request.
func NewDeleteBookingParamsWithContext(ctx context.Context) *DeleteBookingParams {
	return &DeleteBookingParams{
		Context: ctx,
	}
}

// NewDeleteBookingParamsWithHTTPClient creates a new DeleteBookingParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBookingParamsWithHTTPClient(client *http.Client) *DeleteBookingParams {
	return &DeleteBookingParams{
		HTTPClient: client,
	}
}

/*
DeleteBookingParams contains all the parameters to send to the API endpoint

	for the delete booking operation.

	Typically these are written to a http.Request.
*/
type DeleteBookingParams struct {

	/* BookingID.

	   ID of the booking to delete
	*/
	BookingID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete booking params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBookingParams) WithDefaults() *DeleteBookingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete booking params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBookingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete booking params
func (o *DeleteBookingParams) WithTimeout(timeout time.Duration) *DeleteBookingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete booking params
func (o *DeleteBookingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete booking params
func (o *DeleteBookingParams) WithContext(ctx context.Context) *DeleteBookingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete booking params
func (o *DeleteBookingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete booking params
func (o *DeleteBookingParams) WithHTTPClient(client *http.Client) *DeleteBookingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete booking params
func (o *DeleteBookingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBookingID adds the bookingID to the delete booking params
func (o *DeleteBookingParams) WithBookingID(bookingID int64) *DeleteBookingParams {
	o.SetBookingID(bookingID)
	return o
}

// SetBookingID adds the bookingId to the delete booking params
func (o *DeleteBookingParams) SetBookingID(bookingID int64) {
	o.BookingID = bookingID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBookingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bookingID
	if err := r.SetPathParam("bookingID", swag.FormatInt64(o.BookingID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"travel-booking-portal/gen/models"
)

// GetBookingOKCode is the HTTP code returned for type GetBookingOK
const GetBookingOKCode int = 200

/*
GetBookingOK Booking retrieved successfully

swagger:response getBookingOK
*/
type GetBookingOK struct {

	/*
	  In: Body
	*/
	Payload *models.Booking `json:"body,omitempty"`
}

// NewGetBookingOK creates GetBookingOK with default headers values
func NewGetBookingOK() *GetBookingOK {

	return &GetBookingOK{}
}

// WithPayload adds the payload to the get booking o k response
func (o *GetBookingOK) WithPayload(payload *models.Booking) *GetBookingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get booking o k response
func (o *GetBookingOK) SetPayload(payload *models.Booking) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBookingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBookingNotFoundCode is the HTTP code returned for type GetBookingNotFound
const GetBookingNotFoundCode int = 404

/*
GetBookingNotFound Booking not found

swagger:response getBookingNotFound
*/
type GetBookingNotFound struct {
}

// NewGetBookingNotFound creates GetBookingNotFound with default headers values
func NewGetBookingNotFound() *GetBookingNotFound {

	return &GetBookingNotFound{}
}

// WriteResponse to the client
func (o *GetBookingNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetBookingInternalServerErrorCode is the HTTP code returned for type GetBookingInternalServerError
const GetBookingInternalServerErrorCode int = 500

/*
GetBookingInternalServerError Internal Server Error

swagger:response getBookingInternalServerError
*/
type GetBookingInternalServerError struct {
}

// NewGetBookingInternalServerError creates GetBookingInternalServerError with default headers values
func NewGetBookingInternalServerError() *GetBookingInternalServerError {

	return &GetBookingInternalServerError{}
}

// WriteResponse to the client
func (o *GetBookingInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

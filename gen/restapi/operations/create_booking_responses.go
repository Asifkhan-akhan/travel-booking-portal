// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"travel-booking-portal/gen/models"
)

// CreateBookingCreatedCode is the HTTP code returned for type CreateBookingCreated
const CreateBookingCreatedCode int = 201

/*
CreateBookingCreated Booking added successfully

swagger:response createBookingCreated
*/
type CreateBookingCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Booking `json:"body,omitempty"`
}

// NewCreateBookingCreated creates CreateBookingCreated with default headers values
func NewCreateBookingCreated() *CreateBookingCreated {

	return &CreateBookingCreated{}
}

// WithPayload adds the payload to the create booking created response
func (o *CreateBookingCreated) WithPayload(payload *models.Booking) *CreateBookingCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create booking created response
func (o *CreateBookingCreated) SetPayload(payload *models.Booking) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBookingCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBookingConflictCode is the HTTP code returned for type CreateBookingConflict
const CreateBookingConflictCode int = 409

/*
CreateBookingConflict Booking already exists

swagger:response createBookingConflict
*/
type CreateBookingConflict struct {
}

// NewCreateBookingConflict creates CreateBookingConflict with default headers values
func NewCreateBookingConflict() *CreateBookingConflict {

	return &CreateBookingConflict{}
}

// WriteResponse to the client
func (o *CreateBookingConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// CreateBookingInternalServerErrorCode is the HTTP code returned for type CreateBookingInternalServerError
const CreateBookingInternalServerErrorCode int = 500

/*
CreateBookingInternalServerError Internal Server Error

swagger:response createBookingInternalServerError
*/
type CreateBookingInternalServerError struct {
}

// NewCreateBookingInternalServerError creates CreateBookingInternalServerError with default headers values
func NewCreateBookingInternalServerError() *CreateBookingInternalServerError {

	return &CreateBookingInternalServerError{}
}

// WriteResponse to the client
func (o *CreateBookingInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

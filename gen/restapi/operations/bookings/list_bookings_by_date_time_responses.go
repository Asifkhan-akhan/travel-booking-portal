// Code generated by go-swagger; DO NOT EDIT.

package bookings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"travel-booking-portal/gen/models"
)

// ListBookingsByDateTimeOKCode is the HTTP code returned for type ListBookingsByDateTimeOK
const ListBookingsByDateTimeOKCode int = 200

/*
ListBookingsByDateTimeOK List of bookings retrieved successfully

swagger:response listBookingsByDateTimeOK
*/
type ListBookingsByDateTimeOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Booking `json:"body,omitempty"`
}

// NewListBookingsByDateTimeOK creates ListBookingsByDateTimeOK with default headers values
func NewListBookingsByDateTimeOK() *ListBookingsByDateTimeOK {

	return &ListBookingsByDateTimeOK{}
}

// WithPayload adds the payload to the list bookings by date time o k response
func (o *ListBookingsByDateTimeOK) WithPayload(payload []*models.Booking) *ListBookingsByDateTimeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list bookings by date time o k response
func (o *ListBookingsByDateTimeOK) SetPayload(payload []*models.Booking) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListBookingsByDateTimeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Booking, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListBookingsByDateTimeBadRequestCode is the HTTP code returned for type ListBookingsByDateTimeBadRequest
const ListBookingsByDateTimeBadRequestCode int = 400

/*
ListBookingsByDateTimeBadRequest Bad Request - Invalid input parameters

swagger:response listBookingsByDateTimeBadRequest
*/
type ListBookingsByDateTimeBadRequest struct {
}

// NewListBookingsByDateTimeBadRequest creates ListBookingsByDateTimeBadRequest with default headers values
func NewListBookingsByDateTimeBadRequest() *ListBookingsByDateTimeBadRequest {

	return &ListBookingsByDateTimeBadRequest{}
}

// WriteResponse to the client
func (o *ListBookingsByDateTimeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ListBookingsByDateTimeNotFoundCode is the HTTP code returned for type ListBookingsByDateTimeNotFound
const ListBookingsByDateTimeNotFoundCode int = 404

/*
ListBookingsByDateTimeNotFound Bookings not found

swagger:response listBookingsByDateTimeNotFound
*/
type ListBookingsByDateTimeNotFound struct {
}

// NewListBookingsByDateTimeNotFound creates ListBookingsByDateTimeNotFound with default headers values
func NewListBookingsByDateTimeNotFound() *ListBookingsByDateTimeNotFound {

	return &ListBookingsByDateTimeNotFound{}
}

// WriteResponse to the client
func (o *ListBookingsByDateTimeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ListBookingsByDateTimeInternalServerErrorCode is the HTTP code returned for type ListBookingsByDateTimeInternalServerError
const ListBookingsByDateTimeInternalServerErrorCode int = 500

/*
ListBookingsByDateTimeInternalServerError Internal Server Error

swagger:response listBookingsByDateTimeInternalServerError
*/
type ListBookingsByDateTimeInternalServerError struct {
}

// NewListBookingsByDateTimeInternalServerError creates ListBookingsByDateTimeInternalServerError with default headers values
func NewListBookingsByDateTimeInternalServerError() *ListBookingsByDateTimeInternalServerError {

	return &ListBookingsByDateTimeInternalServerError{}
}

// WriteResponse to the client
func (o *ListBookingsByDateTimeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

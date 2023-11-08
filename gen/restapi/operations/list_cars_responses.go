// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"travel-booking-portal/gen/models"
)

// ListCarsOKCode is the HTTP code returned for type ListCarsOK
const ListCarsOKCode int = 200

/*
ListCarsOK List of Cars retrieved successfully

swagger:response listCarsOK
*/
type ListCarsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.CarRental `json:"body,omitempty"`
}

// NewListCarsOK creates ListCarsOK with default headers values
func NewListCarsOK() *ListCarsOK {

	return &ListCarsOK{}
}

// WithPayload adds the payload to the list cars o k response
func (o *ListCarsOK) WithPayload(payload []*models.CarRental) *ListCarsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list cars o k response
func (o *ListCarsOK) SetPayload(payload []*models.CarRental) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCarsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.CarRental, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListCarsNotFoundCode is the HTTP code returned for type ListCarsNotFound
const ListCarsNotFoundCode int = 404

/*
ListCarsNotFound Cars not found

swagger:response listCarsNotFound
*/
type ListCarsNotFound struct {
}

// NewListCarsNotFound creates ListCarsNotFound with default headers values
func NewListCarsNotFound() *ListCarsNotFound {

	return &ListCarsNotFound{}
}

// WriteResponse to the client
func (o *ListCarsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ListCarsInternalServerErrorCode is the HTTP code returned for type ListCarsInternalServerError
const ListCarsInternalServerErrorCode int = 500

/*
ListCarsInternalServerError Internal Server Error

swagger:response listCarsInternalServerError
*/
type ListCarsInternalServerError struct {
}

// NewListCarsInternalServerError creates ListCarsInternalServerError with default headers values
func NewListCarsInternalServerError() *ListCarsInternalServerError {

	return &ListCarsInternalServerError{}
}

// WriteResponse to the client
func (o *ListCarsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
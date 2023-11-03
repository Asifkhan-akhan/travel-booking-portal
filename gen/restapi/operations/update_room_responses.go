// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"travel-booking-portal/gen/models"
)

// UpdateRoomCreatedCode is the HTTP code returned for type UpdateRoomCreated
const UpdateRoomCreatedCode int = 201

/*
UpdateRoomCreated Room updated successfully

swagger:response updateRoomCreated
*/
type UpdateRoomCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Room `json:"body,omitempty"`
}

// NewUpdateRoomCreated creates UpdateRoomCreated with default headers values
func NewUpdateRoomCreated() *UpdateRoomCreated {

	return &UpdateRoomCreated{}
}

// WithPayload adds the payload to the update room created response
func (o *UpdateRoomCreated) WithPayload(payload *models.Room) *UpdateRoomCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update room created response
func (o *UpdateRoomCreated) SetPayload(payload *models.Room) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRoomCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateRoomNotFoundCode is the HTTP code returned for type UpdateRoomNotFound
const UpdateRoomNotFoundCode int = 404

/*
UpdateRoomNotFound Room not found

swagger:response updateRoomNotFound
*/
type UpdateRoomNotFound struct {
}

// NewUpdateRoomNotFound creates UpdateRoomNotFound with default headers values
func NewUpdateRoomNotFound() *UpdateRoomNotFound {

	return &UpdateRoomNotFound{}
}

// WriteResponse to the client
func (o *UpdateRoomNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// UpdateRoomInternalServerErrorCode is the HTTP code returned for type UpdateRoomInternalServerError
const UpdateRoomInternalServerErrorCode int = 500

/*
UpdateRoomInternalServerError Internal Server Error

swagger:response updateRoomInternalServerError
*/
type UpdateRoomInternalServerError struct {
}

// NewUpdateRoomInternalServerError creates UpdateRoomInternalServerError with default headers values
func NewUpdateRoomInternalServerError() *UpdateRoomInternalServerError {

	return &UpdateRoomInternalServerError{}
}

// WriteResponse to the client
func (o *UpdateRoomInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
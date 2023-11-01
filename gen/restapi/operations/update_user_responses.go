// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"travel-booking-portal/gen/models"
)

// UpdateUserCreatedCode is the HTTP code returned for type UpdateUserCreated
const UpdateUserCreatedCode int = 201

/*
UpdateUserCreated User updated successfully

swagger:response updateUserCreated
*/
type UpdateUserCreated struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewUpdateUserCreated creates UpdateUserCreated with default headers values
func NewUpdateUserCreated() *UpdateUserCreated {

	return &UpdateUserCreated{}
}

// WithPayload adds the payload to the update user created response
func (o *UpdateUserCreated) WithPayload(payload *models.User) *UpdateUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user created response
func (o *UpdateUserCreated) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateUserNotFoundCode is the HTTP code returned for type UpdateUserNotFound
const UpdateUserNotFoundCode int = 404

/*
UpdateUserNotFound User already exists

swagger:response updateUserNotFound
*/
type UpdateUserNotFound struct {
}

// NewUpdateUserNotFound creates UpdateUserNotFound with default headers values
func NewUpdateUserNotFound() *UpdateUserNotFound {

	return &UpdateUserNotFound{}
}

// WriteResponse to the client
func (o *UpdateUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// UpdateUserInternalServerErrorCode is the HTTP code returned for type UpdateUserInternalServerError
const UpdateUserInternalServerErrorCode int = 500

/*
UpdateUserInternalServerError Internal Server Error

swagger:response updateUserInternalServerError
*/
type UpdateUserInternalServerError struct {
}

// NewUpdateUserInternalServerError creates UpdateUserInternalServerError with default headers values
func NewUpdateUserInternalServerError() *UpdateUserInternalServerError {

	return &UpdateUserInternalServerError{}
}

// WriteResponse to the client
func (o *UpdateUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

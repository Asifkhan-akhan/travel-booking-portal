// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"travel-booking-portal/gen/models"
)

// CreateUserReader is a Reader for the CreateUser structure.
type CreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewCreateUserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /user] CreateUser", response, response.Code())
	}
}

// NewCreateUserCreated creates a CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {
	return &CreateUserCreated{}
}

/*
CreateUserCreated describes a response with status code 201, with default header values.

User added successfully
*/
type CreateUserCreated struct {
	Payload *models.User
}

// IsSuccess returns true when this create user created response has a 2xx status code
func (o *CreateUserCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user created response has a 3xx status code
func (o *CreateUserCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user created response has a 4xx status code
func (o *CreateUserCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user created response has a 5xx status code
func (o *CreateUserCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create user created response a status code equal to that given
func (o *CreateUserCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create user created response
func (o *CreateUserCreated) Code() int {
	return 201
}

func (o *CreateUserCreated) Error() string {
	return fmt.Sprintf("[POST /user][%d] createUserCreated  %+v", 201, o.Payload)
}

func (o *CreateUserCreated) String() string {
	return fmt.Sprintf("[POST /user][%d] createUserCreated  %+v", 201, o.Payload)
}

func (o *CreateUserCreated) GetPayload() *models.User {
	return o.Payload
}

func (o *CreateUserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserConflict creates a CreateUserConflict with default headers values
func NewCreateUserConflict() *CreateUserConflict {
	return &CreateUserConflict{}
}

/*
CreateUserConflict describes a response with status code 409, with default header values.

User already exists
*/
type CreateUserConflict struct {
}

// IsSuccess returns true when this create user conflict response has a 2xx status code
func (o *CreateUserConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user conflict response has a 3xx status code
func (o *CreateUserConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user conflict response has a 4xx status code
func (o *CreateUserConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user conflict response has a 5xx status code
func (o *CreateUserConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create user conflict response a status code equal to that given
func (o *CreateUserConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create user conflict response
func (o *CreateUserConflict) Code() int {
	return 409
}

func (o *CreateUserConflict) Error() string {
	return fmt.Sprintf("[POST /user][%d] createUserConflict ", 409)
}

func (o *CreateUserConflict) String() string {
	return fmt.Sprintf("[POST /user][%d] createUserConflict ", 409)
}

func (o *CreateUserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserInternalServerError creates a CreateUserInternalServerError with default headers values
func NewCreateUserInternalServerError() *CreateUserInternalServerError {
	return &CreateUserInternalServerError{}
}

/*
CreateUserInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateUserInternalServerError struct {
}

// IsSuccess returns true when this create user internal server error response has a 2xx status code
func (o *CreateUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user internal server error response has a 3xx status code
func (o *CreateUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user internal server error response has a 4xx status code
func (o *CreateUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user internal server error response has a 5xx status code
func (o *CreateUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create user internal server error response a status code equal to that given
func (o *CreateUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create user internal server error response
func (o *CreateUserInternalServerError) Code() int {
	return 500
}

func (o *CreateUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /user][%d] createUserInternalServerError ", 500)
}

func (o *CreateUserInternalServerError) String() string {
	return fmt.Sprintf("[POST /user][%d] createUserInternalServerError ", 500)
}

func (o *CreateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

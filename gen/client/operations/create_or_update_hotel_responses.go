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

// CreateOrUpdateHotelReader is a Reader for the CreateOrUpdateHotel structure.
type CreateOrUpdateHotelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrUpdateHotelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrUpdateHotelCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewCreateOrUpdateHotelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateOrUpdateHotelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /hotel] CreateOrUpdateHotel", response, response.Code())
	}
}

// NewCreateOrUpdateHotelCreated creates a CreateOrUpdateHotelCreated with default headers values
func NewCreateOrUpdateHotelCreated() *CreateOrUpdateHotelCreated {
	return &CreateOrUpdateHotelCreated{}
}

/*
CreateOrUpdateHotelCreated describes a response with status code 201, with default header values.

Hotel added successfully
*/
type CreateOrUpdateHotelCreated struct {
	Payload *models.Hotel
}

// IsSuccess returns true when this create or update hotel created response has a 2xx status code
func (o *CreateOrUpdateHotelCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create or update hotel created response has a 3xx status code
func (o *CreateOrUpdateHotelCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create or update hotel created response has a 4xx status code
func (o *CreateOrUpdateHotelCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create or update hotel created response has a 5xx status code
func (o *CreateOrUpdateHotelCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create or update hotel created response a status code equal to that given
func (o *CreateOrUpdateHotelCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create or update hotel created response
func (o *CreateOrUpdateHotelCreated) Code() int {
	return 201
}

func (o *CreateOrUpdateHotelCreated) Error() string {
	return fmt.Sprintf("[POST /hotel][%d] createOrUpdateHotelCreated  %+v", 201, o.Payload)
}

func (o *CreateOrUpdateHotelCreated) String() string {
	return fmt.Sprintf("[POST /hotel][%d] createOrUpdateHotelCreated  %+v", 201, o.Payload)
}

func (o *CreateOrUpdateHotelCreated) GetPayload() *models.Hotel {
	return o.Payload
}

func (o *CreateOrUpdateHotelCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Hotel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrUpdateHotelConflict creates a CreateOrUpdateHotelConflict with default headers values
func NewCreateOrUpdateHotelConflict() *CreateOrUpdateHotelConflict {
	return &CreateOrUpdateHotelConflict{}
}

/*
CreateOrUpdateHotelConflict describes a response with status code 409, with default header values.

Hotel already exists
*/
type CreateOrUpdateHotelConflict struct {
}

// IsSuccess returns true when this create or update hotel conflict response has a 2xx status code
func (o *CreateOrUpdateHotelConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create or update hotel conflict response has a 3xx status code
func (o *CreateOrUpdateHotelConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create or update hotel conflict response has a 4xx status code
func (o *CreateOrUpdateHotelConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create or update hotel conflict response has a 5xx status code
func (o *CreateOrUpdateHotelConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create or update hotel conflict response a status code equal to that given
func (o *CreateOrUpdateHotelConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create or update hotel conflict response
func (o *CreateOrUpdateHotelConflict) Code() int {
	return 409
}

func (o *CreateOrUpdateHotelConflict) Error() string {
	return fmt.Sprintf("[POST /hotel][%d] createOrUpdateHotelConflict ", 409)
}

func (o *CreateOrUpdateHotelConflict) String() string {
	return fmt.Sprintf("[POST /hotel][%d] createOrUpdateHotelConflict ", 409)
}

func (o *CreateOrUpdateHotelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrUpdateHotelInternalServerError creates a CreateOrUpdateHotelInternalServerError with default headers values
func NewCreateOrUpdateHotelInternalServerError() *CreateOrUpdateHotelInternalServerError {
	return &CreateOrUpdateHotelInternalServerError{}
}

/*
CreateOrUpdateHotelInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateOrUpdateHotelInternalServerError struct {
}

// IsSuccess returns true when this create or update hotel internal server error response has a 2xx status code
func (o *CreateOrUpdateHotelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create or update hotel internal server error response has a 3xx status code
func (o *CreateOrUpdateHotelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create or update hotel internal server error response has a 4xx status code
func (o *CreateOrUpdateHotelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create or update hotel internal server error response has a 5xx status code
func (o *CreateOrUpdateHotelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create or update hotel internal server error response a status code equal to that given
func (o *CreateOrUpdateHotelInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create or update hotel internal server error response
func (o *CreateOrUpdateHotelInternalServerError) Code() int {
	return 500
}

func (o *CreateOrUpdateHotelInternalServerError) Error() string {
	return fmt.Sprintf("[POST /hotel][%d] createOrUpdateHotelInternalServerError ", 500)
}

func (o *CreateOrUpdateHotelInternalServerError) String() string {
	return fmt.Sprintf("[POST /hotel][%d] createOrUpdateHotelInternalServerError ", 500)
}

func (o *CreateOrUpdateHotelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

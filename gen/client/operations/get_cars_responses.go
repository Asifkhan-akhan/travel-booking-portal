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

// GetCarsReader is a Reader for the GetCars structure.
type GetCarsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCarsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCarsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCarsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCarsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /car] GetCars", response, response.Code())
	}
}

// NewGetCarsOK creates a GetCarsOK with default headers values
func NewGetCarsOK() *GetCarsOK {
	return &GetCarsOK{}
}

/*
GetCarsOK describes a response with status code 200, with default header values.

List of Cars retrieved successfully
*/
type GetCarsOK struct {
	Payload []*models.CarRental
}

// IsSuccess returns true when this get cars o k response has a 2xx status code
func (o *GetCarsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cars o k response has a 3xx status code
func (o *GetCarsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cars o k response has a 4xx status code
func (o *GetCarsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cars o k response has a 5xx status code
func (o *GetCarsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cars o k response a status code equal to that given
func (o *GetCarsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cars o k response
func (o *GetCarsOK) Code() int {
	return 200
}

func (o *GetCarsOK) Error() string {
	return fmt.Sprintf("[GET /car][%d] getCarsOK  %+v", 200, o.Payload)
}

func (o *GetCarsOK) String() string {
	return fmt.Sprintf("[GET /car][%d] getCarsOK  %+v", 200, o.Payload)
}

func (o *GetCarsOK) GetPayload() []*models.CarRental {
	return o.Payload
}

func (o *GetCarsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCarsNotFound creates a GetCarsNotFound with default headers values
func NewGetCarsNotFound() *GetCarsNotFound {
	return &GetCarsNotFound{}
}

/*
GetCarsNotFound describes a response with status code 404, with default header values.

Cars not found
*/
type GetCarsNotFound struct {
}

// IsSuccess returns true when this get cars not found response has a 2xx status code
func (o *GetCarsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cars not found response has a 3xx status code
func (o *GetCarsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cars not found response has a 4xx status code
func (o *GetCarsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cars not found response has a 5xx status code
func (o *GetCarsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get cars not found response a status code equal to that given
func (o *GetCarsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get cars not found response
func (o *GetCarsNotFound) Code() int {
	return 404
}

func (o *GetCarsNotFound) Error() string {
	return fmt.Sprintf("[GET /car][%d] getCarsNotFound ", 404)
}

func (o *GetCarsNotFound) String() string {
	return fmt.Sprintf("[GET /car][%d] getCarsNotFound ", 404)
}

func (o *GetCarsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCarsInternalServerError creates a GetCarsInternalServerError with default headers values
func NewGetCarsInternalServerError() *GetCarsInternalServerError {
	return &GetCarsInternalServerError{}
}

/*
GetCarsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCarsInternalServerError struct {
}

// IsSuccess returns true when this get cars internal server error response has a 2xx status code
func (o *GetCarsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cars internal server error response has a 3xx status code
func (o *GetCarsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cars internal server error response has a 4xx status code
func (o *GetCarsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cars internal server error response has a 5xx status code
func (o *GetCarsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get cars internal server error response a status code equal to that given
func (o *GetCarsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get cars internal server error response
func (o *GetCarsInternalServerError) Code() int {
	return 500
}

func (o *GetCarsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /car][%d] getCarsInternalServerError ", 500)
}

func (o *GetCarsInternalServerError) String() string {
	return fmt.Sprintf("[GET /car][%d] getCarsInternalServerError ", 500)
}

func (o *GetCarsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
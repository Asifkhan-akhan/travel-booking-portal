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

// GetBookingReader is a Reader for the GetBooking structure.
type GetBookingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBookingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBookingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetBookingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBookingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /booking/{bookingID}] GetBooking", response, response.Code())
	}
}

// NewGetBookingOK creates a GetBookingOK with default headers values
func NewGetBookingOK() *GetBookingOK {
	return &GetBookingOK{}
}

/*
GetBookingOK describes a response with status code 200, with default header values.

Booking retrieved successfully
*/
type GetBookingOK struct {
	Payload *models.Booking
}

// IsSuccess returns true when this get booking o k response has a 2xx status code
func (o *GetBookingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get booking o k response has a 3xx status code
func (o *GetBookingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get booking o k response has a 4xx status code
func (o *GetBookingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get booking o k response has a 5xx status code
func (o *GetBookingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get booking o k response a status code equal to that given
func (o *GetBookingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get booking o k response
func (o *GetBookingOK) Code() int {
	return 200
}

func (o *GetBookingOK) Error() string {
	return fmt.Sprintf("[GET /booking/{bookingID}][%d] getBookingOK  %+v", 200, o.Payload)
}

func (o *GetBookingOK) String() string {
	return fmt.Sprintf("[GET /booking/{bookingID}][%d] getBookingOK  %+v", 200, o.Payload)
}

func (o *GetBookingOK) GetPayload() *models.Booking {
	return o.Payload
}

func (o *GetBookingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Booking)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBookingNotFound creates a GetBookingNotFound with default headers values
func NewGetBookingNotFound() *GetBookingNotFound {
	return &GetBookingNotFound{}
}

/*
GetBookingNotFound describes a response with status code 404, with default header values.

Booking not found
*/
type GetBookingNotFound struct {
}

// IsSuccess returns true when this get booking not found response has a 2xx status code
func (o *GetBookingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get booking not found response has a 3xx status code
func (o *GetBookingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get booking not found response has a 4xx status code
func (o *GetBookingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get booking not found response has a 5xx status code
func (o *GetBookingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get booking not found response a status code equal to that given
func (o *GetBookingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get booking not found response
func (o *GetBookingNotFound) Code() int {
	return 404
}

func (o *GetBookingNotFound) Error() string {
	return fmt.Sprintf("[GET /booking/{bookingID}][%d] getBookingNotFound ", 404)
}

func (o *GetBookingNotFound) String() string {
	return fmt.Sprintf("[GET /booking/{bookingID}][%d] getBookingNotFound ", 404)
}

func (o *GetBookingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBookingInternalServerError creates a GetBookingInternalServerError with default headers values
func NewGetBookingInternalServerError() *GetBookingInternalServerError {
	return &GetBookingInternalServerError{}
}

/*
GetBookingInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetBookingInternalServerError struct {
}

// IsSuccess returns true when this get booking internal server error response has a 2xx status code
func (o *GetBookingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get booking internal server error response has a 3xx status code
func (o *GetBookingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get booking internal server error response has a 4xx status code
func (o *GetBookingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get booking internal server error response has a 5xx status code
func (o *GetBookingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get booking internal server error response a status code equal to that given
func (o *GetBookingInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get booking internal server error response
func (o *GetBookingInternalServerError) Code() int {
	return 500
}

func (o *GetBookingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /booking/{bookingID}][%d] getBookingInternalServerError ", 500)
}

func (o *GetBookingInternalServerError) String() string {
	return fmt.Sprintf("[GET /booking/{bookingID}][%d] getBookingInternalServerError ", 500)
}

func (o *GetBookingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
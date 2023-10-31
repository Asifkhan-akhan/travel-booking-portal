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

// GetRoomsOfHotelReader is a Reader for the GetRoomsOfHotel structure.
type GetRoomsOfHotelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoomsOfHotelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoomsOfHotelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRoomsOfHotelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoomsOfHotelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /hotel/{ID}/rooms] getRoomsOfHotel", response, response.Code())
	}
}

// NewGetRoomsOfHotelOK creates a GetRoomsOfHotelOK with default headers values
func NewGetRoomsOfHotelOK() *GetRoomsOfHotelOK {
	return &GetRoomsOfHotelOK{}
}

/*
GetRoomsOfHotelOK describes a response with status code 200, with default header values.

rooms in response
*/
type GetRoomsOfHotelOK struct {
	Payload *models.Room
}

// IsSuccess returns true when this get rooms of hotel o k response has a 2xx status code
func (o *GetRoomsOfHotelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get rooms of hotel o k response has a 3xx status code
func (o *GetRoomsOfHotelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rooms of hotel o k response has a 4xx status code
func (o *GetRoomsOfHotelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get rooms of hotel o k response has a 5xx status code
func (o *GetRoomsOfHotelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get rooms of hotel o k response a status code equal to that given
func (o *GetRoomsOfHotelOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get rooms of hotel o k response
func (o *GetRoomsOfHotelOK) Code() int {
	return 200
}

func (o *GetRoomsOfHotelOK) Error() string {
	return fmt.Sprintf("[GET /hotel/{ID}/rooms][%d] getRoomsOfHotelOK  %+v", 200, o.Payload)
}

func (o *GetRoomsOfHotelOK) String() string {
	return fmt.Sprintf("[GET /hotel/{ID}/rooms][%d] getRoomsOfHotelOK  %+v", 200, o.Payload)
}

func (o *GetRoomsOfHotelOK) GetPayload() *models.Room {
	return o.Payload
}

func (o *GetRoomsOfHotelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Room)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoomsOfHotelNotFound creates a GetRoomsOfHotelNotFound with default headers values
func NewGetRoomsOfHotelNotFound() *GetRoomsOfHotelNotFound {
	return &GetRoomsOfHotelNotFound{}
}

/*
GetRoomsOfHotelNotFound describes a response with status code 404, with default header values.

Rooms not found
*/
type GetRoomsOfHotelNotFound struct {
}

// IsSuccess returns true when this get rooms of hotel not found response has a 2xx status code
func (o *GetRoomsOfHotelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rooms of hotel not found response has a 3xx status code
func (o *GetRoomsOfHotelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rooms of hotel not found response has a 4xx status code
func (o *GetRoomsOfHotelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get rooms of hotel not found response has a 5xx status code
func (o *GetRoomsOfHotelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get rooms of hotel not found response a status code equal to that given
func (o *GetRoomsOfHotelNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get rooms of hotel not found response
func (o *GetRoomsOfHotelNotFound) Code() int {
	return 404
}

func (o *GetRoomsOfHotelNotFound) Error() string {
	return fmt.Sprintf("[GET /hotel/{ID}/rooms][%d] getRoomsOfHotelNotFound ", 404)
}

func (o *GetRoomsOfHotelNotFound) String() string {
	return fmt.Sprintf("[GET /hotel/{ID}/rooms][%d] getRoomsOfHotelNotFound ", 404)
}

func (o *GetRoomsOfHotelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoomsOfHotelInternalServerError creates a GetRoomsOfHotelInternalServerError with default headers values
func NewGetRoomsOfHotelInternalServerError() *GetRoomsOfHotelInternalServerError {
	return &GetRoomsOfHotelInternalServerError{}
}

/*
GetRoomsOfHotelInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetRoomsOfHotelInternalServerError struct {
}

// IsSuccess returns true when this get rooms of hotel internal server error response has a 2xx status code
func (o *GetRoomsOfHotelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get rooms of hotel internal server error response has a 3xx status code
func (o *GetRoomsOfHotelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get rooms of hotel internal server error response has a 4xx status code
func (o *GetRoomsOfHotelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get rooms of hotel internal server error response has a 5xx status code
func (o *GetRoomsOfHotelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get rooms of hotel internal server error response a status code equal to that given
func (o *GetRoomsOfHotelInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get rooms of hotel internal server error response
func (o *GetRoomsOfHotelInternalServerError) Code() int {
	return 500
}

func (o *GetRoomsOfHotelInternalServerError) Error() string {
	return fmt.Sprintf("[GET /hotel/{ID}/rooms][%d] getRoomsOfHotelInternalServerError ", 500)
}

func (o *GetRoomsOfHotelInternalServerError) String() string {
	return fmt.Sprintf("[GET /hotel/{ID}/rooms][%d] getRoomsOfHotelInternalServerError ", 500)
}

func (o *GetRoomsOfHotelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

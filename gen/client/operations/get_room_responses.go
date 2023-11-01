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

// GetRoomReader is a Reader for the GetRoom structure.
type GetRoomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRoomNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoomInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /room/{roomID}] GetRoom", response, response.Code())
	}
}

// NewGetRoomOK creates a GetRoomOK with default headers values
func NewGetRoomOK() *GetRoomOK {
	return &GetRoomOK{}
}

/*
GetRoomOK describes a response with status code 200, with default header values.

Room retrieved successfully
*/
type GetRoomOK struct {
	Payload *models.Room
}

// IsSuccess returns true when this get room o k response has a 2xx status code
func (o *GetRoomOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get room o k response has a 3xx status code
func (o *GetRoomOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get room o k response has a 4xx status code
func (o *GetRoomOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get room o k response has a 5xx status code
func (o *GetRoomOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get room o k response a status code equal to that given
func (o *GetRoomOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get room o k response
func (o *GetRoomOK) Code() int {
	return 200
}

func (o *GetRoomOK) Error() string {
	return fmt.Sprintf("[GET /room/{roomID}][%d] getRoomOK  %+v", 200, o.Payload)
}

func (o *GetRoomOK) String() string {
	return fmt.Sprintf("[GET /room/{roomID}][%d] getRoomOK  %+v", 200, o.Payload)
}

func (o *GetRoomOK) GetPayload() *models.Room {
	return o.Payload
}

func (o *GetRoomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Room)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoomNotFound creates a GetRoomNotFound with default headers values
func NewGetRoomNotFound() *GetRoomNotFound {
	return &GetRoomNotFound{}
}

/*
GetRoomNotFound describes a response with status code 404, with default header values.

Room not found
*/
type GetRoomNotFound struct {
}

// IsSuccess returns true when this get room not found response has a 2xx status code
func (o *GetRoomNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get room not found response has a 3xx status code
func (o *GetRoomNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get room not found response has a 4xx status code
func (o *GetRoomNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get room not found response has a 5xx status code
func (o *GetRoomNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get room not found response a status code equal to that given
func (o *GetRoomNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get room not found response
func (o *GetRoomNotFound) Code() int {
	return 404
}

func (o *GetRoomNotFound) Error() string {
	return fmt.Sprintf("[GET /room/{roomID}][%d] getRoomNotFound ", 404)
}

func (o *GetRoomNotFound) String() string {
	return fmt.Sprintf("[GET /room/{roomID}][%d] getRoomNotFound ", 404)
}

func (o *GetRoomNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoomInternalServerError creates a GetRoomInternalServerError with default headers values
func NewGetRoomInternalServerError() *GetRoomInternalServerError {
	return &GetRoomInternalServerError{}
}

/*
GetRoomInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRoomInternalServerError struct {
}

// IsSuccess returns true when this get room internal server error response has a 2xx status code
func (o *GetRoomInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get room internal server error response has a 3xx status code
func (o *GetRoomInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get room internal server error response has a 4xx status code
func (o *GetRoomInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get room internal server error response has a 5xx status code
func (o *GetRoomInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get room internal server error response a status code equal to that given
func (o *GetRoomInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get room internal server error response
func (o *GetRoomInternalServerError) Code() int {
	return 500
}

func (o *GetRoomInternalServerError) Error() string {
	return fmt.Sprintf("[GET /room/{roomID}][%d] getRoomInternalServerError ", 500)
}

func (o *GetRoomInternalServerError) String() string {
	return fmt.Sprintf("[GET /room/{roomID}][%d] getRoomInternalServerError ", 500)
}

func (o *GetRoomInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Room room
//
// swagger:model Room
type Room struct {

	// hotel id
	HotelID int64 `json:"hotel_id,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// number
	Number string `json:"number,omitempty"`
}

// Validate validates this room
func (m *Room) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this room based on context it is used
func (m *Room) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Room) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Room) UnmarshalBinary(b []byte) error {
	var res Room
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

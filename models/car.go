package models

import (
	"github.com/fatih/structs"
)

type CarRental struct {
	ID     int  `json:"id" db:"id" structs:"id" bson:"_id"`
	Booked bool `json:"booked" bson:"booked" db:"booked" structs:"booked"`
}

// Map converts structs to a map representation
func (s *CarRental) Map() map[string]interface{} {
	return structs.Map(s)
}

// Names returns the field names of CarRental model
func (s *CarRental) Names() []string {
	fields := structs.Fields(s)
	names := make([]string, len(fields))

	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}

package models

import (
	"github.com/fatih/structs"
	"time"
)

// Booking of a hotel room or a car.
type Booking struct {
	ID          int       `json:"id" bson:"id" db:"id" structs:"id"`
	UserID      int       `json:"user_id" bson:"user_id" db:"user_id" structs:"user_id"`
	ServiceType string    `json:"service_type" bson:"service_type" db:"service_type" structs:"service_type"` // Either Hotel or Car
	ServiceID   int       `json:"service_id" bson:"service_id" db:"service_id" structs:"service_id"`
	FromDate    time.Time `json:"from_date" bson:"from_date" db:"from_date" structs:"from_date"`
	ToDate      time.Time `json:"to_date" bson:"to_date" db:"to_date" structs:"to_date"`
	Confirmed   bool      `json:"confirmed" bson:"confirmed" db:"confirmed" structs:"confirmed"`
	Penalty     bool      `json:"penalty" bson:"penalty" db:"penalty" structs:"penalty"`
	Paid        bool      `json:"paid" bson:"paid" db:"paid" structs:"paid"`
}

// Map converts structs to a map representation
func (s *Booking) Map() map[string]interface{} {
	return structs.Map(s)
}

// Names returns the field names of Booking model
func (s *Booking) Names() []string {
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

package models

import (
	"reflect"
	"testing"
	"time"
)

func TestBooking_Map(t *testing.T) {
	type fields struct {
		ID          int
		UserID      int
		ServiceType string
		ServiceID   int

		FromDate  time.Time
		ToDate    time.Time
		Confirmed bool
		Penalty   bool
		Paid      bool
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "passed - booking map",
			fields: fields{
				ID:          12,
				UserID:      1,
				ServiceType: "CarRental",
				ServiceID:   2,
				FromDate:    time.Date(2023, 10, 23, 12, 0, 0, 0, time.UTC), // Create a valid time.Time value
				ToDate:      time.Date(2023, 10, 23, 13, 0, 0, 0, time.UTC),
				Confirmed:   true,
				Penalty:     false,
				Paid:        true,
			},
			want: map[string]interface{}{
				"id":           12,
				"user_id":      1,
				"service_type": "CarRental",
				"service_id":   2,
				"from_date":    time.Date(2023, 10, 23, 12, 0, 0, 0, time.UTC), // Create a valid time.Time value
				"to_date":      time.Date(2023, 10, 23, 13, 0, 0, 0, time.UTC),
				"confirmed":    true,
				"penalty":      false,
				"paid":         true,
			},
		},
		{
			name: "success - student struct to map with fewer fields",
			fields: fields{
				ID:          1,
				UserID:      101,
				ServiceType: "CarRental",
				ServiceID:   2,
				FromDate:    time.Date(2023, 10, 23, 12, 0, 0, 0, time.UTC), // Create a valid time.Time value
				ToDate:      time.Date(2023, 10, 23, 13, 0, 0, 0, time.UTC),
				Confirmed:   true,
				Penalty:     false,
				Paid:        true,
			},
			want: map[string]interface{}{
				"id":           1,
				"user_id":      101,
				"service_type": "CarRental",
				"service_id":   2,
				"from_date":    time.Date(2023, 10, 23, 12, 0, 0, 0, time.UTC), // Create a valid time.Time value
				"to_date":      time.Date(2023, 10, 23, 13, 0, 0, 0, time.UTC),
				"confirmed":    true,
				"penalty":      false,
				"paid":         true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booking := Booking{
				ID:          tt.fields.ID,
				UserID:      tt.fields.UserID,
				ServiceType: tt.fields.ServiceType,
				ServiceID:   tt.fields.ServiceID,
				FromDate:    tt.fields.FromDate,
				ToDate:      tt.fields.ToDate,
				Confirmed:   tt.fields.Confirmed,
				Penalty:     tt.fields.Penalty,
				Paid:        tt.fields.Paid,
			}
			if got := booking.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBooking_Names(t *testing.T) {
	type fields struct {
		ID          int
		UserID      int
		ServiceType string
		ServiceID   int
		Date        time.Time
		Confirmed   bool
		Penalty     bool
		Paid        bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "success - names of booking struct fields",
			fields: fields{},
			want:   []string{"id", "user_id", "service_type", "service_id", "from_date", "to_date", "confirmed", "penalty", "paid"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booking := Booking{
				ID:          tt.fields.ID,
				UserID:      tt.fields.UserID,
				ServiceType: tt.fields.ServiceType,
				ServiceID:   tt.fields.ServiceID,
				FromDate:    tt.fields.Date,
				ToDate:      tt.fields.Date,
				Confirmed:   tt.fields.Confirmed,
				Penalty:     tt.fields.Penalty,
				Paid:        tt.fields.Paid,
			}
			if got := booking.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}

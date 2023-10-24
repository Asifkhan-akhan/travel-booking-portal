package models

import (
	"reflect"
	"testing"
)

func TestCarRental_Map(t *testing.T) {
	type fields struct {
		ID     int
		Booked bool
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success - car rental struct to map",
			fields: fields{
				ID:     12,
				Booked: false,
			},
			want: map[string]interface{}{
				"id":     12,
				"booked": false,
			},
		},
		{
			name: "success - car rental struct to map with fewer fields",
			fields: fields{
				Booked: false,
			},
			want: map[string]interface{}{
				"id":     0, // You can change the expected value as needed.
				"booked": false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CarRental{
				ID:     tt.fields.ID,
				Booked: tt.fields.Booked,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarRental_Names(t *testing.T) {
	type fields struct {
		ID     int
		Booked bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "success - names of car rental struct fields",
			fields: fields{
				ID:     1,
				Booked: false,
			},
			want: []string{"id", "booked"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CarRental{
				ID:     tt.fields.ID,
				Booked: tt.fields.Booked,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}

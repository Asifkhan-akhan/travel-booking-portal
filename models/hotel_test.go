package models

import (
	"reflect"
	"testing"
)

func TestHotel_Map(t *testing.T) {
	type fields struct {
		ID          int
		Name        string
		Location    string
		Description string
		Rooms       int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success - hotel struct to map",
			fields: fields{
				ID:          1,
				Name:        "peshawar Hotel",
				Location:    "peshawar",
				Description: "A luxurious hotel",
				Rooms:       100,
			},
			want: map[string]interface{}{
				"id":          1,
				"name":        "peshawar Hotel",
				"location":    "peshawar",
				"description": "A luxurious hotel",
				"rooms":       100,
			},
		},
		{
			name: "success - hotel struct to map with fewer fields",
			fields: fields{
				Name: "peshawar Hotel",
			},
			want: map[string]interface{}{
				"id":          0, // You can change the expected value as needed.
				"name":        "peshawar Hotel",
				"location":    "",
				"description": "",
				"rooms":       0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Hotel{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				Location:    tt.fields.Location,
				Description: tt.fields.Description,
				Rooms:       tt.fields.Rooms,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHotel_Names(t *testing.T) {
	type fields struct {
		ID          int
		Name        string
		Location    string
		Description string
		Rooms       int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "success - names of hotel struct fields",
			fields: fields{
				ID:          1,
				Name:        "abc Hotel",
				Location:    "City Center",
				Description: "A luxurious hotel",
				Rooms:       100,
			},
			want: []string{"id", "name", "location", "description", "rooms"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Hotel{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				Location:    tt.fields.Location,
				Description: tt.fields.Description,
				Rooms:       tt.fields.Rooms,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoom_Map(t *testing.T) {
	type fields struct {
		ID      int
		Number  string
		HotelID int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success - room struct to map",
			fields: fields{
				ID:      1,
				Number:  "101",
				HotelID: 2,
			},
			want: map[string]interface{}{
				"id":       1,
				"number":   "101",
				"hotel_id": 2,
			},
		},
		{
			name: "success - room struct to map with fewer fields",
			fields: fields{
				Number: "102",
			},
			want: map[string]interface{}{
				"id":       0, // You can change the expected value as needed.
				"number":   "102",
				"hotel_id": 0, // You can change the expected value as needed.
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Room{
				ID:      tt.fields.ID,
				Number:  tt.fields.Number,
				HotelID: tt.fields.HotelID,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoom_Names(t *testing.T) {
	type fields struct {
		ID      int
		Number  string
		HotelID int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "success - names of room struct fields",
			fields: fields{
				ID:      1,
				Number:  "101",
				HotelID: 2,
			},
			want: []string{"id", "number", "hotel_id"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Room{
				ID:      tt.fields.ID,
				Number:  tt.fields.Number,
				HotelID: tt.fields.HotelID,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}

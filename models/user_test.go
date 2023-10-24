package models

import (
	"reflect"
	"testing"
)

func TestUser_Map(t *testing.T) {
	type fields struct {
		ID    int
		Name  string
		Email string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success - user struct to map",
			fields: fields{
				ID:    1,
				Name:  "khan",
				Email: "asif@example.com",
			},
			want: map[string]interface{}{
				"id":    1,
				"name":  "khan",
				"email": "asif@example.com",
			},
		},
		{
			name: "success - user struct to map with fewer fields",
			fields: fields{
				Name: "asif",
			},
			want: map[string]interface{}{
				"id":    0, // You can change the expected value as needed.
				"name":  "asif",
				"email": "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &User{
				ID:    tt.fields.ID,
				Name:  tt.fields.Name,
				Email: tt.fields.Email,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Names(t *testing.T) {
	type fields struct {
		ID    int
		Name  string
		Email string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "success - names of user struct fields",
			fields: fields{
				ID:    1,
				Name:  "khan",
				Email: "khan@example.com",
			},
			want: []string{"id", "name", "email"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &User{
				ID:    tt.fields.ID,
				Name:  tt.fields.Name,
				Email: tt.fields.Email,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}

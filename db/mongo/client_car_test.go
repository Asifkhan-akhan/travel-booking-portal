package mongo

import (
	"os"
	"testing"
	"travel-booking-portal/models"
)

func Test_client_CreateOrUpdateCarRental(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost") // Corrected DB_HOST value
	type args struct {
		car *models.CarRental
	}

	tests := []struct {
		name    string
		args    args // Use *models.Booking here
		wantErr bool
	}{
		{
			name: "success - carrental added in db",
			args: args{car: &models.CarRental{
				ID:     23,
				Booked: false,
			}},
			wantErr: false,
		},
		{
			name: "fail - fail to add car rental in db",
			args: args{car: &models.CarRental{
				ID:     22,
				Booked: true,
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient()
			_, err := c.CreateOrUpdateCarRental(tt.args.car)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCarRental() error = %v, wantErr %v", err, tt.wantErr)
				t.Errorf("this is serr")
			}
		})
	}
}

func Test_client_DeleteCarRental(t *testing.T) {
	// Set up the MongoDB client and collection for testing
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")
	c, _ := NewClient()
	car := &models.CarRental{
		ID:     1,
		Booked: true,
	}
	// Create a CarRental to be deleted
	createdCarID, createCarErr := c.CreateOrUpdateCarRental(car)

	if createCarErr != nil {
		t.Errorf("CreateCarRental() error: %v", createCarErr)
	} else {
		t.Logf("Car rental created with ID: %d", createdCarID)
	}

	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - car rental deleted from db",
			args:    args{id: createdCarID}, // Use the ID of the created car rental
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteCarRental(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCarRental() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func Test_client_ListCarRentals(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")
	c, _ := NewClient()

	// Create a sample car rental
	carRental := &models.CarRental{
		ID:     1,
		Booked: false,
	}

	// Create a single car rental
	createdCarRentalID, createCarRentalErr := c.CreateOrUpdateCarRental(carRental)
	if createCarRentalErr != nil {
		t.Errorf("CreateCarRental() error: %v", createCarRentalErr)
	} else {
		t.Logf("Car rental created with ID: %d", createdCarRentalID)
	}

	tests := []struct {
		name           string
		searchCriteria map[string]interface{}
		want           []*models.CarRental
		wantErr        bool
	}{
		{
			name:           "success - all car rentals get",
			searchCriteria: map[string]interface{}{},       // No search criteria
			want:           []*models.CarRental{carRental}, // You can modify this as needed
			wantErr:        false,
		},
		{
			name:           "success - getting a single car rental",
			searchCriteria: map[string]interface{}{"id": createdCarRentalID}, // Search by ID
			want:           []*models.CarRental{carRental},                   // You can modify this as needed
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			carRentals, err := c.ListCarRental(tt.searchCriteria)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListCarRentals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Got %v", carRentals)
		})
	}
}

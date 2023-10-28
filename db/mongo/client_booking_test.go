package mongo

import (
	"os"
	"testing"
	"time"
	"travel-booking-portal/models"
)

func Test_client_CreateOrUpdateBooking(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost") // Corrected DB_HOST value
	type args struct {
		booking *models.Booking
	}

	tests := []struct {
		name    string
		args    args // Use *models.Booking here
		wantErr bool
	}{
		{
			name: "success - booking added in db",
			args: args{booking: &models.Booking{
				ID:          2,
				UserID:      1,
				ServiceType: "Hotel",

				ServiceID: 122223,
				FromDate:  time.Now(),
				Confirmed: true,
				Penalty:   false,
				Paid:      true,
			}},
			wantErr: false,
		},
		{
			name: "fail - booking added in db",
			args: args{booking: &models.Booking{
				ID:          11,
				ServiceType: "Car",
				ServiceID:   456,
				FromDate:    time.Now(),
				Confirmed:   false,
				Penalty:     true,
				Paid:        false,
			}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient()
			_, err := c.CreateOrUpdateBooking(tt.args.booking)
			if (err != nil) != tt.wantErr {

				t.Errorf("CreateBooking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_DeleteBooking(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient()
	booking := &models.Booking{
		UserID:      1,
		ServiceType: "CarRental",

		ServiceID: 123,
		FromDate:  time.Now(),
		Confirmed: true,
		Penalty:   false,
		Paid:      true,
	}
	// _, _ = c.CreateBooking(booking)
	createdBookingID, createBookingErr := c.CreateOrUpdateBooking(booking)

	if createBookingErr != nil {
		t.Logf("CreateBooking() error: %v", createBookingErr)
	} else {
		t.Logf("Booking created with ID: %d", createdBookingID)
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
			name:    "success - booking deleted from db",
			args:    args{id: booking.ID},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteBooking(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteBooking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_GetBooking(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")
	c, _ := NewClient()
	booking := &models.Booking{
		ID:          3232,
		UserID:      121212,
		ServiceType: "CarRental",
		ServiceID:   123,
		Confirmed:   true,
		Penalty:     false,
		Paid:        true,
	}

	// Create a single booking
	createdBookingID, createBookingErr := c.CreateOrUpdateBooking(booking)
	if createBookingErr != nil {
		t.Errorf("CreateBooking() error: %v", createBookingErr)
	} else {
		t.Logf("Booking created with ID: %d", createdBookingID)
	}

	type args struct {
		id int
	}

	tests := []struct {
		name    string
		args    args
		want    []*models.Booking
		wantErr bool
	}{

		{
			name:    "success - all bookings get",
			args:    args{id: 0},
			want:    []*models.Booking{booking},
			wantErr: false,
		},
		{
			name:    "success -  getting sing document",
			args:    args{id: createdBookingID},
			want:    []*models.Booking{booking},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booking, err := c.GetBooking(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBooking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Got %v", booking)
		})
	}

}

func TestSearchBooking(t *testing.T) {
	// Set up a test MongoDB client (replace with your actual setup).
	c, _ := NewClient()

	// Create some sample bookings with different user IDs and dates.
	booking1 := &models.Booking{
		ID:          1,
		UserID:      101,
		ServiceType: "Hotel",
		ServiceID:   1,
		FromDate:    time.Date(2023, time.October, 26, 0, 0, 0, 0, time.UTC),
		Confirmed:   true,
		Penalty:     false,
		Paid:        true,
	}
	booking2 := &models.Booking{
		ID:          2,
		UserID:      102,
		ServiceType: "Hotel",
		ServiceID:   2,
		FromDate:    time.Date(2023, time.October, 27, 0, 0, 0, 0, time.UTC),
		Confirmed:   true,
		Penalty:     false,
		Paid:        true,
	}
	booking3 := &models.Booking{
		ID:          3,
		UserID:      101,
		ServiceType: "Hotel",
		ServiceID:   3,
		FromDate:    time.Date(2023, time.October, 26, 0, 0, 0, 0, time.UTC),
		Confirmed:   true,
		Penalty:     false,
		Paid:        true,
	}

	// Create bookings in the database (replace with your actual creation logic).
	c.CreateOrUpdateBooking(booking1)
	c.CreateOrUpdateBooking(booking2)
	c.CreateOrUpdateBooking(booking3)

	// Search for bookings for a specific user and date.
	userID := 101
	bookings, err := c.SearchBooking(userID, "2023, 10, 26")
	if err != nil {
		t.Errorf("Error searching for bookings: %v", err)
	}

	// Log the results.
	for _, booking := range bookings {
		t.Logf("Booking ID: %d, User ID: %d, Date: %s", booking.ID, booking.UserID, booking.FromDate)
	}
}

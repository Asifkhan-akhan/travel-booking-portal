package mongo

import (
	"os"
	"testing"
	"travel-booking-portal/models"
)

func Test_client_CreateOrUpdateHotel(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost") // Corrected DB_HOST value
	type args struct {
		hotel *models.Hotel
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - hotel added in db",
			args: args{hotel: &models.Hotel{

				Name:        "marhab",
				Location:    "karachi",
				Description: "A beautiful hotel",
				Rooms:       100,
			}},
			wantErr: false,
		},
		{
			name: "fail - fail to add hotel in db",
			args: args{hotel: &models.Hotel{
				ID:          0,
				Name:        "Invalid Hotel",
				Location:    "Invalid Location",
				Description: "An invalid hotel",
				Rooms:       0, // Invalid number of rooms
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient()
			_, err := c.CreateOrUpdateHotel(tt.args.hotel)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateHotel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_DeleteHotel(t *testing.T) {
	// Set up the MongoDB client and collection for testing
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")
	c, _ := NewClient()
	hotel := &models.Hotel{
		ID:          1,
		Name:        "Test Hotel",
		Location:    "Test Location",
		Description: "Test Description",
		Rooms:       10,
	}
	// Create a Hotel to be deleted
	createdHotelID, createHotelErr := c.CreateOrUpdateHotel(hotel)

	if createHotelErr != nil {
		t.Errorf("CreateHotel() error: %v", createHotelErr)
	} else {
		t.Logf("Hotel created with ID: %d", createdHotelID)
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
			name:    "success - hotel deleted from db",
			args:    args{id: createdHotelID}, // Use the ID of the created hotel
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteHotel(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteHotel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func Test_client_GetHotels(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")
	c, _ := NewClient()

	// Create a sample hotel
	hotel := &models.Hotel{
		ID:          1,
		Name:        "Sample Hotel",
		Location:    "Sample Location",
		Description: "A sample hotel description",
		Rooms:       50,
	}

	// Create a single hotel
	createdHotelID, createHotelErr := c.CreateOrUpdateHotel(hotel)
	if createHotelErr != nil {
		t.Errorf("CreateHotel() error: %v", createHotelErr)
	} else {
		t.Logf("Hotel created with ID: %d", createdHotelID)
	}

	type args struct {
		id int
	}

	tests := []struct {
		name    string
		args    args
		want    []*models.Hotel
		wantErr bool
	}{
		{
			name:    "success - all hotels get",
			args:    args{id: 0},
			want:    []*models.Hotel{hotel}, // You can modify this as needed
			wantErr: false,
		},
		{
			name:    "success - getting a single hotel",
			args:    args{id: createdHotelID},
			want:    []*models.Hotel{hotel}, // You can modify this as needed
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hotels, err := c.GetHotel(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHotels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Got %v", hotels)
		})
	}
}
func Test_client_CreateOrUpdateRoom(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost") // Corrected DB_HOST value
	type args struct {
		room *models.Room
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - room added in db",
			args: args{room: &models.Room{
				Number:  "101",
				HotelID: 1,
			}},
			wantErr: false,
		},
		{
			name: "fail - fail to add room in db",
			args: args{room: &models.Room{
				Number:  "", // Invalid room number
				HotelID: 0,  // Invalid hotel ID
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient()
			_, err := c.CreateOrUpdateRoom(tt.args.room)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrUpdateRoom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_DeleteRoom(t *testing.T) {
	// Set up the MongoDB client and collection for testing
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")
	c, _ := NewClient()
	room := &models.Room{
		ID:      1,
		Number:  "101",
		HotelID: 1,
	}
	// Create a Room to be deleted
	createdRoomID, createRoomErr := c.CreateOrUpdateRoom(room)

	if createRoomErr != nil {
		t.Errorf("CreateRoom() error: %v", createRoomErr)
	} else {
		t.Logf("Room created with ID: %d", createdRoomID)
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
			name:    "success - room deleted from db",
			args:    args{id: createdRoomID}, // Use the ID of the created room
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteRoom(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteRoom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_GetRooms(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")
	c, _ := NewClient()

	// Create a sample room
	room := &models.Room{
		ID:      1,
		Number:  "101",
		HotelID: 123,
	}

	// Create a single room
	createdRoomID, createRoomErr := c.CreateOrUpdateRoom(room)
	if createRoomErr != nil {
		t.Errorf("CreateRoom() error: %v", createRoomErr)
	} else {
		t.Logf("Room created with ID: %d", createdRoomID)
	}

	type args struct {
		id int
	}

	tests := []struct {
		name    string
		args    args
		want    []*models.Room
		wantErr bool
	}{
		{
			name:    "success - all rooms get",
			args:    args{id: 0},
			want:    []*models.Room{room}, // You can modify this as needed
			wantErr: false,
		},
		{
			name:    "success - getting a single room",
			args:    args{id: createdRoomID},
			want:    []*models.Room{room}, // You can modify this as needed
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rooms, err := c.GetRoom(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRooms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Got %v", rooms)
		})
	}
}

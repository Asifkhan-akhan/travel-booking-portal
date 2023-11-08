package mongo

import (
	"os"
	"reflect"
	"testing"
	"travel-booking-portal/models"
)

func Test_client_CreatepdateUser(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost") // Corrected DB_HOST value
	type args struct {
		user *models.User
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - room added in db",
			args: args{user: &models.User{

				Name:  "khan",
				Email: "wancloud@gmail.com",
			}},
			wantErr: false,
		},
		{
			name: "fail - fail to add room in db",
			args: args{user: &models.User{
				Name:  "",    // Invalid room number
				Email: "abs", // Invalid hotel ID
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient()
			_, err := c.CreateOrUpdateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrUpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_DeleteUser(t *testing.T) {
	// Set up the MongoDB client and collection for testing
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")
	c, _ := NewClient()
	user := &models.User{
		ID:    1,
		Name:  "ahmad",
		Email: "ali@example.com",
	}
	// Create a User to be deleted
	createdUserID, createUserErr := c.CreateOrUpdateUser(user)

	if createUserErr != nil {
		t.Errorf("CreateOrUpdateUser() error: %v", createUserErr)
	} else {
		t.Logf("User created with ID: %d", createdUserID)
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
			name:    "success - user deleted from db",
			args:    args{id: createdUserID}, // Use the ID of the created user
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteUser(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func Test_client_ListUser(t *testing.T) {
	// Set up the MongoDB client and collection for testing
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")
	c, _ := NewClient()
	user := &models.User{
		ID:    1,
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}
	createdUserID, createUserErr := c.CreateOrUpdateUser(user)

	if createUserErr != nil {
		t.Errorf("CreateOrUpdateUser() error: %v", createUserErr)
	} else {
		t.Logf("User created with ID: %d", createdUserID)
	}

	tests := []struct {
		name           string
		searchCriteria map[string]interface{}
		want           []*models.User
		wantErr        bool
	}{
		{
			name:           "success - user get",
			searchCriteria: map[string]interface{}{"_id": createdUserID}, // Search by ID
			want:           []*models.User{user},
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.ListUser(tt.searchCriteria)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

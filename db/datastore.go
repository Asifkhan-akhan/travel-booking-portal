package db

import (
	"travel-booking-portal/models"
)

// DataStore defines the interface for CRUD operations on models.

type DataStore interface {

	// User CRUD operations
	CreateOrUpdateUser(user *models.User) (int, error)
	ListUser(searchCriteria map[string]interface{}) ([]*models.User, error)
	DeleteUser(userID int) error

	// Hotel CRUD operations
	CreateOrUpdateHotel(hotel *models.Hotel) (int, error)
	ListHotel(searchCriteria map[string]interface{}) ([]*models.Hotel, error)
	DeleteHotel(hotelID int) error

	// CarRental CRUD operations
	CreateOrUpdateCarRental(car *models.CarRental) (int, error)
	ListCarRental(searchCriteria map[string]interface{}) ([]*models.CarRental, error)
	DeleteCarRental(carID int) error

	// Room CRUD operations
	CreateOrUpdateRoom(room *models.Room) (int, error)
	ListRoom(searchCriteria map[string]interface{}) ([]*models.Room, error)
	DeleteRoom(roomID int) error

	// Booking CRUD operations
	CreateOrUpdateBooking(booking *models.Booking) (int, error)
	ListBooking(searchCriteria map[string]interface{}) ([]*models.Booking, error)
	DeleteBooking(bookingID int) error
}

type Options struct {
	TestMode bool
}

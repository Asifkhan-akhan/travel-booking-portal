package db

import (
	"travel-booking-portal/models"
)

// DataStore defines the interface for CRUD operations on models.

type DataStore interface {

	// User CRUD operations
	CreateOrUpdateUser(user *models.User) (int, error)
	GetUser(userID int) ([]*models.User, error)
	DeleteUser(userID int) error

	// Hotel CRUD operations
	CreateOrUpdateHotel(hotel *models.Hotel) (int, error)
	GetHotel(hotelID int) ([]*models.Hotel, error)
	DeleteHotel(hotelID int) error

	// CarRental CRUD operations
	CreateOrUpdateCarRental(car *models.CarRental) (int, error)
	GetCarRental(carID int) ([]*models.CarRental, error)
	DeleteCarRental(carID int) error

	// Room CRUD operations
	CreateOrUpdateRoom(room *models.Room) (int, error)
	GetRoom(roomID int) ([]*models.Room, error)
	DeleteRoom(roomID int) error

	// Booking CRUD operations
	CreateOrUpdateBooking(booking *models.Booking) (int, error)
	GetBooking(bookingID int) ([]*models.Booking, error)
	DeleteBooking(bookingID int) error
	SearchBooking(id int, date string) ([]*models.Booking, error)
	ConfirmedBooking(id int) ([]*models.Booking, error)
}

type Options struct {
	TestMode bool
}

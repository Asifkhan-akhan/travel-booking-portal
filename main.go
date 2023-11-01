package main

import (
	"fmt"
	"time"
	"travel-booking-portal/db/mongo"
	"travel-booking-portal/models"
)

// this is for temporary  testing

func main() {
	// Initialize your MongoDB client.
	c, _ := mongo.NewClient()

	booking2 := &models.Booking{
		ID:          5,
		UserID:      102,
		ServiceType: "CarRental",
		ServiceID:   3,
		FromDate:    time.Date(2001, time.October, 22, 0, 0, 0, 0, time.UTC),
		ToDate:      time.Date(2001, time.October, 26, 0, 0, 0, 0, time.UTC),

		Confirmed: true,
		Penalty:   false,
		Paid:      true,
	}
	booking3 := &models.Booking{
		ID:          1112,
		UserID:      22222,
		ServiceType: "CarRental",
		ServiceID:   13,
		FromDate:    time.Date(2001, time.October, 4, 0, 0, 0, 0, time.UTC),
		ToDate:      time.Date(2001, time.October, 5, 0, 0, 0, 0, time.UTC),
		Confirmed:   true,
		Penalty:     true,
		Paid:        true,
	}

	// Create bookings in the database (replace with your actual creation logic).
	c.CreateOrUpdateBooking(booking2)
	c.CreateOrUpdateBooking(booking3)

	// Search for bookings for a specific user and date.
	userID := 101
	//2023-10-26T00:00:00.000+00:00
	searchDate := "2023, 10, 03"
	bookings, err := c.SearchBooking(userID, searchDate)
	if err != nil {
		fmt.Println("Error searching for bookings: ", err)
	}

	// Log the results.
	for _, booking := range bookings {
		fmt.Println("Booking ID: ", booking.ID, "User ID: ", booking.UserID, "Date: ", booking.FromDate)
	}
}

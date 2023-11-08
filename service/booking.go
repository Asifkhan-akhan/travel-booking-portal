package service

import (
	"fmt"
	"net/smtp"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"travel-booking-portal/config"
	"travel-booking-portal/models"
)

func sendEmail(receiverEmail string) {
	auth := smtp.PlainAuth("", viper.GetString(config.EmailSMTPUsername), viper.GetString(config.EmailSMTPPassword), viper.GetString(config.EmailSMTPHost))
	subject := "Reminder for Your Booking"
	body := fmt.Sprintf(`Dear User,
This is a reminder for your upcoming %s (ID: %d). Your service is scheduled to start in 5 minutes.
Thank you for choosing our services.

`)
	message := []byte("Subject: " + subject + "\r\n" + body)

	err := smtp.SendMail(fmt.Sprintf("%s:%s", viper.GetString(config.EmailSMTPHost), viper.GetString(config.EmailSMTPPort)), auth, viper.GetString(config.EmailFromAddress), []string{receiverEmail}, message)
	if err != nil {
		fmt.Println("Failed to send email:", err)
	} else {
		fmt.Println("Email sent successfully.")
	}
}

func (s *Service) CreateBooking(booking *models.Booking) (int, error) {
	if booking.ToDate.Before(booking.FromDate) || booking.ToDate.Equal(booking.FromDate) || booking.ToDate.Before(time.Now()) {
		return 0, errors.New("Incorrect times are given")
	}
	bookingID, err := s.db.CreateOrUpdateBooking(booking)
	if err != nil {
		return 0, err
	}

	emailTime := booking.FromDate.Add(-5 * time.Minute)
	durationToWait := emailTime.Sub(time.Now())

	if durationToWait > 0 {

		users, err := s.db.ListUser(map[string]interface{}{"_id": booking.UserID})
		if err != nil {
			return 0, err
		}
		if len(users) > 0 {
			user := users[0]
			go scheduleEmail(user.Email, durationToWait)
		} else {
			return 0, errors.New("No given user find")
		}
	}

	return bookingID, nil
}

func scheduleEmail(email string, durationToWait time.Duration) {
	time.Sleep(durationToWait)
	sendEmail(email)
}

func (s *Service) UpdateBooking(booking *models.Booking) (int, error) {

	existingBookings, _ := s.GetBooking(booking.ID)
	if len(existingBookings) == 0 {
		return 0, errors.New("Booking not found")
	}

	existingBooking := existingBookings[0]

	// Check if the booking is confirmed
	if existingBooking.Confirmed {
		return 0, errors.New("Cannot edit a confirmed booking")
	}

	timeDifference := existingBooking.FromDate.Sub(time.Now())

	if timeDifference < 48*time.Hour {

		return 0, errors.New("Cannot edit a booking less than 48 hours before the booking date")
	}
	if booking.ToDate.Before(booking.FromDate) || booking.ToDate.Equal(booking.FromDate) || booking.ToDate.Before(time.Now()) {
		return 0, errors.New("Incorrect times are given")
	}
	bookingID, err := s.db.CreateOrUpdateBooking(booking)
	if err != nil {
		return 0, err
	}

	return bookingID, nil
}

func (s *Service) DeleteBooking(bookingID int) error {
	bookings, err := s.GetBooking(bookingID)
	if err != nil {
		return err
	}

	if len(bookings) == 0 {
		return errors.New("Booking not found")
	}

	booking := bookings[0]

	if booking.Confirmed {
		return errors.New("Cannot delete a confirmed booking")
	}

	timeDifference := booking.FromDate.Sub(time.Now())

	if timeDifference < 48*time.Hour {
		return errors.New("Cannot delete a booking less than 48 hours before the booking time")
	}

	return s.db.DeleteBooking(bookingID)
}

// GetBookingByID retrieves a booking by its ID.
func (s *Service) GetBooking(bookingID int) ([]*models.Booking, error) {
	// Call the data layer method to retrieve the booking by ID.
	booking, err := s.db.ListBooking(map[string]interface{}{"_id": bookingID})
	if err != nil {

		return nil, err
	}

	// Check if the booking was found.
	if len(booking) > 0 {
		return booking, nil
	}

	return nil, err
}

// ListAllBookings retrieves a list of all bookings.
func (s *Service) ListBookings() ([]*models.Booking, error) {

	bookings, err := s.db.ListBooking(nil) // Pass nil to indicate no specific filter criteria.
	if err != nil {

		return nil, err
	}

	return bookings, nil
}

// GetBookingsByDateTime retrieves bookings for a user based on a datetime.
func (s *Service) GetBookingsByDateTime(userID int, dateTime time.Time) ([]*models.Booking, error) {

	searchParams := map[string]interface{}{
		"user_id":   userID,
		"from_date": dateTime,
	}

	bookings, err := s.db.ListBooking(searchParams)
	if err != nil {
		return nil, err
	}

	return bookings, nil
}
func (s *Service) ListConfirmedBookings(userID int) ([]*models.Booking, error) {
	searchParams := map[string]interface{}{
		"user_id":   userID,
		"confirmed": true,
	}

	// Call the GetBooking method from the data layer with the search parameter.
	bookings, err := s.db.ListBooking(searchParams)
	if err != nil {
		// Handle any errors that may occur.
		return nil, err
	}

	// Return the filtered bookings.
	return bookings, nil
}

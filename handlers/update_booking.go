package handlers

import (
	"time"

	"github.com/go-openapi/runtime/middleware"

	runtime "travel-booking-portal"
	"travel-booking-portal/gen/restapi/operations"
	"travel-booking-portal/models"
)

// NewEditStudent function is for edit student
func NewUpdateBooking(rt *runtime.Runtime) operations.UpdateBookingHandler {
	return &updateBooking{rt: rt}
}

type updateBooking struct {
	rt *runtime.Runtime
}

// Handle the put student request
func (d *updateBooking) Handle(params operations.UpdateBookingParams) middleware.Responder {

	booking := &models.Booking{
		ID:          int(params.BookingID),
		UserID:      int(params.Booking.UserID),
		ServiceType: params.Booking.ServiceType,
		ServiceID:   int(params.Booking.ServiceID),
		FromDate:    time.Time(params.Booking.FromDate),
		ToDate:      time.Time(params.Booking.ToDate),
		Confirmed:   params.Booking.Confirmed,
		Penalty:     params.Booking.Penalty,
		Paid:        params.Booking.Paid,
	}
	if _, err := d.rt.Service().UpdateBooking(booking); err != nil {
		if err.Error() == "Booking not found" {
			log().Errorf(err.Error())
			return operations.NewDeleteBookingNotFound()
		}
		if err.Error() == "Cannot edit a confirmed booking" {
			log().Errorf(err.Error())
			return operations.NewDeleteBookingBadRequest()
		}
		if err.Error() == "Cannot edit a booking less than 48 hours before the booking service start" {
			log().Errorf(err.Error())
			return operations.NewDeleteBookingBadRequest()
		}
		log().Errorf(err.Error())
		return operations.NewUpdateBookingInternalServerError()
	}

	return operations.NewUpdateBookingOK()
}

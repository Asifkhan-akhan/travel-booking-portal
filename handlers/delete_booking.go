package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "travel-booking-portal"

	"travel-booking-portal/gen/restapi/operations"
)

func NewDeleteBooking(rt *runtime.Runtime) operations.DeleteBookingHandler {
	return &deleteBooking{rt: rt}
}

type deleteBooking struct {
	rt *runtime.Runtime
}

// Handle the delete entry request
func (d *deleteBooking) Handle(params operations.DeleteBookingParams) middleware.Responder {
	if err := d.rt.Service().DeleteBooking(int(params.BookingID)); err != nil {
		if err.Error() == "Booking not found" {
			log().Errorf(err.Error())
			return operations.NewDeleteBookingNotFound()
		}
		if err.Error() == "Cannot delete a confirmed booking" {
			log().Errorf(err.Error())
			return operations.NewDeleteBookingBadRequest()
		}
		if err.Error() == "Cannot delete a booking less than 48 hours before the booking time" {
			log().Errorf(err.Error())
			return operations.NewDeleteBookingBadRequest()
		}
		log().Errorf(err.Error())
		return operations.NewDeleteBookingInternalServerError()
	}

	return operations.NewDeleteBookingNoContent()
}

package handlers

import (
	"time"
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/restapi/operations"
	"travel-booking-portal/models"

	"github.com/go-openapi/runtime/middleware"
)

// NewCreateUser handles requests for creating a user
func NewCreateBooking(rt *runtime.Runtime) operations.CreateBookingHandler {
	return &createBookingHandler{rt: rt}
}

type createBookingHandler struct {
	rt *runtime.Runtime
}

// Handle the create user request
func (d *createBookingHandler) Handle(params operations.CreateBookingParams) middleware.Responder {
	id, err := d.rt.Service().CreateBooking(&models.Booking{
		UserID:      int(params.Booking.UserID),
		ServiceType: params.Booking.ServiceType,
		ServiceID:   int(params.Booking.ServiceID),
		FromDate:    time.Time(params.Booking.FromDate),
		ToDate:      time.Time(params.Booking.ToDate),
		Confirmed:   params.Booking.Confirmed,
		Penalty:     params.Booking.Penalty,
		Paid:        params.Booking.Paid,
	})
	if err != nil {
		if err.Error() == "Incorrect times are given" {

			log().Errorf("Error %s", err)
			return operations.NewCreateBookingBadRequest()
		}
		if err.Error() == "No given user find" {

			log().Errorf("Error %s", err)
			return operations.NewCreateBookingBadRequest()
		}

		log().Errorf("failed to create booking: %s", err)
	}
	if id == 0 {
		log().Errorf("Error : Conflict found in give request data ")
		return &operations.CreateBookingConflict{}
	}
	params.Booking.ID = int64(id)
	return operations.NewCreateBookingCreated().WithPayload(params.Booking)
}

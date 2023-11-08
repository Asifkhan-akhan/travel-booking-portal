package handlers

import (
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/models"
	"travel-booking-portal/gen/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewListBookingHandler handles a request for retrieving a user
func NewListBooking(rt *runtime.Runtime) operations.ListBookingsHandler {
	return &listBookingHandler{rt: rt}
}

type listBookingHandler struct {
	rt *runtime.Runtime
}

func (d *listBookingHandler) Handle(params operations.ListBookingsParams) middleware.Responder {
	users, err := d.rt.Service().ListBookings()
	if err != nil {
		return operations.NewListBookingsInternalServerError()
	}

	var swaggerBookings []*models.Booking

	for _, booking := range users {
		swaggerBookings = append(swaggerBookings, &models.Booking{
			ID:          int64(booking.ID),
			UserID:      int64(booking.UserID),
			ServiceID:   int64(booking.ServiceID),
			ServiceType: booking.ServiceType,
			FromDate:    strfmt.DateTime(booking.FromDate),
			ToDate:      strfmt.DateTime(booking.ToDate),
			Penalty:     booking.Penalty,
			Confirmed:   booking.Confirmed,
			Paid:        booking.Paid,
		})
	}

	if len(swaggerBookings) > 0 {
		return operations.NewListBookingsOK().WithPayload(swaggerBookings)
	}

	return operations.NewListBookingsNotFound()
}

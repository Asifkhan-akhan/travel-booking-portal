package handlers

import (
	"time"
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/models"
	"travel-booking-portal/gen/restapi/operations/bookings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewListConfirmedBookingHandler handles a request for retrieving a user
func NewListBookingByDateTime(rt *runtime.Runtime) bookings.ListBookingsByDateTimeHandler {
	return &listBookingbyDateTimeHandler{rt: rt}
}

type listBookingbyDateTimeHandler struct {
	rt *runtime.Runtime
}

func (d *listBookingbyDateTimeHandler) Handle(params bookings.ListBookingsByDateTimeParams) middleware.Responder {
	userID := int(params.UserID)
	datetime := params.DateTime

	// Parse the string into a time.Time value.
	parseddateTime, err := time.Parse(time.RFC3339Nano, datetime)
	users, err := d.rt.Service().GetBookingsByDateTime(userID, parseddateTime)
	if err != nil {
		return bookings.NewListConfirmedBookingsInternalServerError()
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
		return bookings.NewListConfirmedBookingsOK().WithPayload(swaggerBookings)
	}

	return bookings.NewListConfirmedBookingsNotFound()
}

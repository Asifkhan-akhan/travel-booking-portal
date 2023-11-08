package handlers

import (
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/models"
	"travel-booking-portal/gen/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func NewGetBooking(rt *runtime.Runtime) operations.GetBookingHandler {
	return &getBookingHandler{rt: rt}
}

type getBookingHandler struct {
	rt *runtime.Runtime
}

// Handle the get user request
func (d *getBookingHandler) Handle(params operations.GetBookingParams) middleware.Responder {
	bookings, err := d.rt.Service().GetBooking(int(params.BookingID))

	if err != nil {
		return operations.NewGetBookingInternalServerError()
	}

	if len(bookings) > 0 {
		swaggerbooking := &models.Booking{

			UserID:      int64(bookings[0].UserID),
			ServiceID:   int64(bookings[0].ServiceID),
			ServiceType: bookings[0].ServiceType,
			FromDate:    strfmt.DateTime(bookings[0].FromDate),
			ToDate:      strfmt.DateTime(bookings[0].ToDate),
			Confirmed:   bookings[0].Confirmed,
			Penalty:     bookings[0].Penalty,
			Paid:        bookings[0].Paid,
		}

		return operations.NewGetBookingOK().WithPayload(swaggerbooking)

	} else {

		return operations.NewGetBookingNotFound()

	}
}

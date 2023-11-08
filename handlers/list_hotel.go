package handlers

import (
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/models"
	"travel-booking-portal/gen/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

// NewListHotelHandler handles a request for retrieving a user
func NewListHotel(rt *runtime.Runtime) operations.ListHotelsHandler {
	return &listHotelHandler{rt: rt}
}

type listHotelHandler struct {
	rt *runtime.Runtime
}

func (d *listHotelHandler) Handle(params operations.ListHotelsParams) middleware.Responder {
	users, err := d.rt.Service().ListHotels()
	if err != nil {
		return operations.NewListHotelsInternalServerError()
	}

	var swaggerHotels []*models.Hotel

	for _, hotel := range users {
		swaggerHotels = append(swaggerHotels, &models.Hotel{
			ID:          int64(hotel.ID),
			Name:        hotel.Name,
			Location:    hotel.Location,
			Description: hotel.Description,
			Rooms:       int64(hotel.Rooms),
		})
	}

	if len(swaggerHotels) > 0 {
		return operations.NewListHotelsOK().WithPayload(swaggerHotels)
	}

	return operations.NewListHotelsNotFound()
}

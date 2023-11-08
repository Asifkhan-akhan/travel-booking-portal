package handlers

import (
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/restapi/operations"
	"travel-booking-portal/models"

	"github.com/go-openapi/runtime/middleware"
)

// NewCreateUser handles requests for creating a user
func NewCreateHotel(rt *runtime.Runtime) operations.CreateHotelHandler {
	return &createHotelHandler{rt: rt}
}

type createHotelHandler struct {
	rt *runtime.Runtime
}

// Handle the create user request
func (d *createHotelHandler) Handle(params operations.CreateHotelParams) middleware.Responder {
	id, err := d.rt.Service().CreateHotel(&models.Hotel{
		Name:        params.Hotel.Name,
		Location:    params.Hotel.Location,
		Description: params.Hotel.Description,
		Rooms:       int(params.Hotel.Rooms),
	})

	if err != nil {
		log().Errorf("failed to create hotel: %s", err)
	}

	params.Hotel.ID = int64(id)
	return operations.NewCreateHotelCreated().WithPayload(params.Hotel)
}

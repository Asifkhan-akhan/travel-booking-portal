package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "travel-booking-portal"
	"travel-booking-portal/gen/restapi/operations"
	"travel-booking-portal/models"
)

// NewEditStudent function is for edit student
func NewUpdateHotl(rt *runtime.Runtime) operations.UpdateHotelHandler {
	return &updateHotel{rt: rt}
}

type updateHotel struct {
	rt *runtime.Runtime
}

// Handle the put student request
func (d *updateHotel) Handle(params operations.UpdateHotelParams) middleware.Responder {

	hotel := &models.Hotel{
		ID:          int(params.HotelID),
		Name:        params.Hotel.Name,
		Location:    params.Hotel.Location,
		Description: params.Hotel.Description,
		Rooms:       int(params.Hotel.Rooms),
	}
	if _, err := d.rt.Service().UpdateHotel(hotel); err != nil {
		return operations.NewUpdateHotelInternalServerError()
	}

	return operations.NewUpdateHotelOK()
}

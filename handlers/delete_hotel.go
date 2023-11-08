package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "travel-booking-portal"

	"travel-booking-portal/gen/restapi/operations"
)

func NewDeleteHotel(rt *runtime.Runtime) operations.DeleteHotelHandler {
	return &deleteHotel{rt: rt}
}

type deleteHotel struct {
	rt *runtime.Runtime
}

// Handle the delete entry request
func (d *deleteHotel) Handle(params operations.DeleteHotelParams) middleware.Responder {
	if err := d.rt.Service().DeleteHotel(int(params.HotelID)); err != nil {

		return operations.NewDeleteHotelInternalServerError()
	}

	return operations.NewDeleteHotelNoContent()
}

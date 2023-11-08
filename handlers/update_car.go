package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "travel-booking-portal"
	"travel-booking-portal/gen/restapi/operations"
	"travel-booking-portal/models"
)

// NewEditStudent function is for edit student
func NewUpdateCar(rt *runtime.Runtime) operations.UpdateCarHandler {
	return &updateCar{rt: rt}
}

type updateCar struct {
	rt *runtime.Runtime
}

// Handle the put student request
func (d *updateCar) Handle(params operations.UpdateCarParams) middleware.Responder {

	car := &models.CarRental{
		ID:     int(params.Car.ID),
		Booked: params.Car.Booked,
	}
	if _, err := d.rt.Service().UpdateCarRental(car); err != nil {
		return operations.NewUpdateCarInternalServerError()
	}

	return operations.NewUpdateCarOK()
}

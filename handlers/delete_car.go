package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "travel-booking-portal"

	"travel-booking-portal/gen/restapi/operations"
)

func NewDeleteCar(rt *runtime.Runtime) operations.DeleteCarHandler {
	return &deleteCar{rt: rt}
}

type deleteCar struct {
	rt *runtime.Runtime
}

// Handle the delete entry request
func (d *deleteCar) Handle(params operations.DeleteCarParams) middleware.Responder {
	if err := d.rt.Service().DeleteCarRental(int(params.CarID)); err != nil {

		return operations.NewDeleteCarInternalServerError()
	}

	return operations.NewDeleteCarNoContent()
}

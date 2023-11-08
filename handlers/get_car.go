package handlers

import (
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/models"
	"travel-booking-portal/gen/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func NewGetCar(rt *runtime.Runtime) operations.GetCarHandler {
	return &getCarHandler{rt: rt}
}

type getCarHandler struct {
	rt *runtime.Runtime
}

// Handle the get user request
func (d *getCarHandler) Handle(params operations.GetCarParams) middleware.Responder {
	cars, err := d.rt.Service().GetCar(int(params.CarID))

	if err != nil {
		return operations.NewGetCarInternalServerError()
	}

	if len(cars) > 0 {
		swaggercar := &models.CarRental{
			Booked: cars[0].Booked,
			ID:     int64(cars[0].ID),
		}

		return operations.NewGetCarOK().WithPayload(swaggercar)

	} else {

		return operations.NewGetCarNotFound()

	}
}

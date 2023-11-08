package handlers

import (
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/models"
	"travel-booking-portal/gen/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

// NewListCarHandler handles a request for retrieving a user
func NewListCar(rt *runtime.Runtime) operations.ListCarsHandler {
	return &listCarHandler{rt: rt}
}

type listCarHandler struct {
	rt *runtime.Runtime
}

func (d *listCarHandler) Handle(params operations.ListCarsParams) middleware.Responder {
	users, err := d.rt.Service().ListCars()
	if err != nil {
		return operations.NewListCarsInternalServerError()
	}

	var swaggerCars []*models.CarRental

	for _, car := range users {
		swaggerCars = append(swaggerCars, &models.CarRental{
			ID:     int64(car.ID),
			Booked: car.Booked,
		})
	}

	if len(swaggerCars) > 0 {
		return operations.NewListCarsOK().WithPayload(swaggerCars)
	}

	return operations.NewListCarsNotFound()
}

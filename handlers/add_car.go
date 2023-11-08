package handlers

import (
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/restapi/operations"
	"travel-booking-portal/models"

	"github.com/go-openapi/runtime/middleware"
)

// NewCreateUser handles requests for creating a car
func NewCreateCar(rt *runtime.Runtime) operations.CreateCarHandler {
	return &createCarHandler{rt: rt}

}

type createCarHandler struct {
	rt *runtime.Runtime
}

// Handle the create  request
func (d *createCarHandler) Handle(params operations.CreateCarParams) middleware.Responder {
	id, err := d.rt.Service().CreateCarRental(&models.CarRental{
		Booked: params.Car.Booked,
	})

	if err != nil {
		log().Errorf("failed to create Car: %s", err)
	}

	params.Car.ID = int64(id)
	return operations.NewCreateCarCreated().WithPayload(params.Car)
}

package handlers

import (
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/restapi/operations"
	"travel-booking-portal/models"

	"github.com/go-openapi/runtime/middleware"
)

// NewCreateUser handles requests for creating a room
func NewCreateRoom(rt *runtime.Runtime) operations.CreateRoomHandler {
	return &createRoomHandler{rt: rt}
}

type createRoomHandler struct {
	rt *runtime.Runtime
}

// Handle the create room request
func (d *createRoomHandler) Handle(params operations.CreateRoomParams) middleware.Responder {
	id, err := d.rt.Service().CreateRoom(&models.Room{
		Number:  params.Room.Number,
		HotelID: int(params.Room.HotelID),
	})

	if err != nil {
		log().Errorf("failed to create Room: %s", err)
		if err.Error() == "Failed to fetch the hotel" {
			return operations.NewUpdateRoomBadRequest()
		}
	}
	if id == 0 {
		return operations.NewCreateRoomBadRequest()
	}

	params.Room.ID = int64(id)
	return operations.NewCreateRoomCreated().WithPayload(params.Room)
}

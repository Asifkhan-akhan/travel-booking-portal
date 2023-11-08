package handlers

import (
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/models"
	"travel-booking-portal/gen/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func NewGetRoom(rt *runtime.Runtime) operations.GetRoomHandler {
	return &getRoomHandler{rt: rt}
}

type getRoomHandler struct {
	rt *runtime.Runtime
}

// Handle the get user request
func (d *getRoomHandler) Handle(params operations.GetRoomParams) middleware.Responder {
	rooms, err := d.rt.Service().GetRoom(int(params.RoomID))

	if err != nil {
		return operations.NewGetRoomInternalServerError()
	}
	if len(rooms) > 0 {
		swaggerroom := &models.Room{
			Number:  rooms[0].Number,
			HotelID: int64(rooms[0].HotelID),
		}

		return operations.NewGetRoomOK().WithPayload(swaggerroom)
	} else {

		// Handle the case when no user is found
		return operations.NewGetRoomNotFound()
	}
}

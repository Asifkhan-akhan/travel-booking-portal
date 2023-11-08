package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "travel-booking-portal"

	"travel-booking-portal/gen/restapi/operations"
)

func NewDeleteRoom(rt *runtime.Runtime) operations.DeleteRoomHandler {
	return &deleteRoom{rt: rt}
}

type deleteRoom struct {
	rt *runtime.Runtime
}

// Handle the delete entry request
func (d *deleteRoom) Handle(params operations.DeleteRoomParams) middleware.Responder {
	if err := d.rt.Service().DeleteRoom(int(params.RoomID)); err != nil {

		return operations.NewDeleteRoomInternalServerError()
	}

	return operations.NewDeleteRoomNoContent()
}

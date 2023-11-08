package handlers

import (
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/models"
	"travel-booking-portal/gen/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

// NewListRoomHandler handles a request for retrieving a user
func NewListRoom(rt *runtime.Runtime) operations.ListRoomsHandler {
	return &listRoomHandler{rt: rt}
}

type listRoomHandler struct {
	rt *runtime.Runtime
}

func (d *listRoomHandler) Handle(params operations.ListRoomsParams) middleware.Responder {
	users, err := d.rt.Service().ListRooms()
	if err != nil {
		return operations.NewListRoomsInternalServerError()
	}

	var swaggerRooms []*models.Room

	for _, room := range users {
		swaggerRooms = append(swaggerRooms, &models.Room{
			ID:      int64(room.ID),
			Number:  room.Number,
			HotelID: int64(room.HotelID),
		})
	}

	if len(swaggerRooms) > 0 {
		return operations.NewListRoomsOK().WithPayload(swaggerRooms)
	}

	return operations.NewListRoomsNotFound()
}

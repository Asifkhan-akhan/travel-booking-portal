package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "travel-booking-portal"
	"travel-booking-portal/gen/restapi/operations"
	"travel-booking-portal/models"
)

// NewEditStudent function is for edit student
func NewUpdateRoom(rt *runtime.Runtime) operations.UpdateRoomHandler {
	return &updateRoom{rt: rt}
}

type updateRoom struct {
	rt *runtime.Runtime
}

// Handle the put student request
func (d *updateRoom) Handle(params operations.UpdateRoomParams) middleware.Responder {

	room := &models.Room{
		ID:      int(params.Room.ID),
		Number:  params.Room.Number,
		HotelID: int(params.Room.HotelID),
	}
	if _, err := d.rt.Service().UpdateRoom(room); err != nil {
		log().Errorf("failed to create Room: %s", err)
		if err.Error() == "Failed to fetch the hotel" {
			return operations.NewUpdateRoomBadRequest()
		}
		return operations.NewUpdateRoomInternalServerError()
	}

	return operations.NewUpdateRoomOK()
}

package handlers

import (
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/models"
	"travel-booking-portal/gen/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func NewGetHotel(rt *runtime.Runtime) operations.GetHotelHandler {
	return &getHotelHandler{rt: rt}
}

type getHotelHandler struct {
	rt *runtime.Runtime
}

func (d *getHotelHandler) Handle(params operations.GetHotelParams) middleware.Responder {
	hotels, err := d.rt.Service().GetHotel(int(params.HotelID))

	if err != nil {
		return operations.NewGetHotelInternalServerError()
	}

	if len(hotels) > 0 {
		swaggerhotel := &models.Hotel{
			Name:        hotels[0].Name,
			Location:    hotels[0].Location,
			Description: hotels[0].Description,
			Rooms:       int64(hotels[0].Rooms),
		}

		return operations.NewGetHotelOK().WithPayload(swaggerhotel)

	} else {

		return operations.NewGetHotelNotFound()

	}
}

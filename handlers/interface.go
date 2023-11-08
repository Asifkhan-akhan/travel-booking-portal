package handlers

import (
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/restapi/operations"

	"github.com/go-openapi/loads"
)

// Handler replaces swagger handler
type Handler *operations.TravelBookingPortalAPI

// NewHandler overrides swagger api handlers
func NewHandler(rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewTravelBookingPortalAPI(spec)

	// User handlers
	handler.CreateUserHandler = NewCreateUser(rt)
	handler.GetUserHandler = NewGetUser(rt)
	handler.UpdateUserHandler = NewUpdateUser(rt)
	handler.DeleteUserHandler = NewDeleteUser(rt)
	handler.ListUsersHandler = NewListUser(rt)

	// Car handlers
	handler.CreateCarHandler = NewCreateCar(rt)
	handler.GetCarHandler = NewGetCar(rt)
	handler.UpdateCarHandler = NewUpdateCar(rt)
	handler.DeleteCarHandler = NewDeleteCar(rt)
	handler.ListCarsHandler = NewListCar(rt)

	// Hotel handlers
	handler.CreateHotelHandler = NewCreateHotel(rt)
	handler.GetHotelHandler = NewGetHotel(rt)
	handler.UpdateHotelHandler = NewUpdateHotl(rt)
	handler.DeleteHotelHandler = NewDeleteHotel(rt)
	handler.ListHotelsHandler = NewListHotel(rt)

	// Room handlers
	handler.CreateRoomHandler = NewCreateRoom(rt)
	handler.GetRoomHandler = NewGetRoom(rt)
	handler.UpdateRoomHandler = NewUpdateRoom(rt)
	handler.DeleteRoomHandler = NewDeleteRoom(rt)
	handler.ListRoomsHandler = NewListRoom(rt)

	// Booking handlers
	handler.CreateBookingHandler = NewCreateBooking(rt)
	handler.GetBookingHandler = NewGetBooking(rt)
	handler.UpdateBookingHandler = NewUpdateBooking(rt)
	handler.DeleteBookingHandler = NewDeleteBooking(rt)
	handler.ListBookingsHandler = NewListBooking(rt)
	handler.BookingsListConfirmedBookingsHandler = NewListConfirmedBooking(rt)
	handler.BookingsListBookingsByDateTimeHandler = NewListBookingByDateTime(rt)

	return handler
}

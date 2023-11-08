// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"travel-booking-portal/gen/restapi/operations"
	"travel-booking-portal/gen/restapi/operations/bookings"
)

//go:generate swagger generate server --target ../../gen --name TravelBookingPortal --spec ../../swagger.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.TravelBookingPortalAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TravelBookingPortalAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CreateBookingHandler == nil {
		api.CreateBookingHandler = operations.CreateBookingHandlerFunc(func(params operations.CreateBookingParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateBooking has not yet been implemented")
		})
	}
	if api.CreateCarHandler == nil {
		api.CreateCarHandler = operations.CreateCarHandlerFunc(func(params operations.CreateCarParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateCar has not yet been implemented")
		})
	}
	if api.CreateHotelHandler == nil {
		api.CreateHotelHandler = operations.CreateHotelHandlerFunc(func(params operations.CreateHotelParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateHotel has not yet been implemented")
		})
	}
	if api.CreateRoomHandler == nil {
		api.CreateRoomHandler = operations.CreateRoomHandlerFunc(func(params operations.CreateRoomParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateRoom has not yet been implemented")
		})
	}
	if api.CreateUserHandler == nil {
		api.CreateUserHandler = operations.CreateUserHandlerFunc(func(params operations.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateUser has not yet been implemented")
		})
	}
	if api.DeleteBookingHandler == nil {
		api.DeleteBookingHandler = operations.DeleteBookingHandlerFunc(func(params operations.DeleteBookingParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteBooking has not yet been implemented")
		})
	}
	if api.DeleteCarHandler == nil {
		api.DeleteCarHandler = operations.DeleteCarHandlerFunc(func(params operations.DeleteCarParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteCar has not yet been implemented")
		})
	}
	if api.DeleteHotelHandler == nil {
		api.DeleteHotelHandler = operations.DeleteHotelHandlerFunc(func(params operations.DeleteHotelParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteHotel has not yet been implemented")
		})
	}
	if api.DeleteRoomHandler == nil {
		api.DeleteRoomHandler = operations.DeleteRoomHandlerFunc(func(params operations.DeleteRoomParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteRoom has not yet been implemented")
		})
	}
	if api.DeleteUserHandler == nil {
		api.DeleteUserHandler = operations.DeleteUserHandlerFunc(func(params operations.DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteUser has not yet been implemented")
		})
	}
	if api.GetBookingHandler == nil {
		api.GetBookingHandler = operations.GetBookingHandlerFunc(func(params operations.GetBookingParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetBooking has not yet been implemented")
		})
	}
	if api.GetCarHandler == nil {
		api.GetCarHandler = operations.GetCarHandlerFunc(func(params operations.GetCarParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetCar has not yet been implemented")
		})
	}
	if api.GetHotelHandler == nil {
		api.GetHotelHandler = operations.GetHotelHandlerFunc(func(params operations.GetHotelParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetHotel has not yet been implemented")
		})
	}
	if api.GetRoomHandler == nil {
		api.GetRoomHandler = operations.GetRoomHandlerFunc(func(params operations.GetRoomParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetRoom has not yet been implemented")
		})
	}
	if api.GetUserHandler == nil {
		api.GetUserHandler = operations.GetUserHandlerFunc(func(params operations.GetUserParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetUser has not yet been implemented")
		})
	}
	if api.ListBookingsHandler == nil {
		api.ListBookingsHandler = operations.ListBookingsHandlerFunc(func(params operations.ListBookingsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListBookings has not yet been implemented")
		})
	}
	if api.BookingsListBookingsByDateTimeHandler == nil {
		api.BookingsListBookingsByDateTimeHandler = bookings.ListBookingsByDateTimeHandlerFunc(func(params bookings.ListBookingsByDateTimeParams) middleware.Responder {
			return middleware.NotImplemented("operation bookings.ListBookingsByDateTime has not yet been implemented")
		})
	}
	if api.ListCarsHandler == nil {
		api.ListCarsHandler = operations.ListCarsHandlerFunc(func(params operations.ListCarsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListCars has not yet been implemented")
		})
	}
	if api.BookingsListConfirmedBookingsHandler == nil {
		api.BookingsListConfirmedBookingsHandler = bookings.ListConfirmedBookingsHandlerFunc(func(params bookings.ListConfirmedBookingsParams) middleware.Responder {
			return middleware.NotImplemented("operation bookings.ListConfirmedBookings has not yet been implemented")
		})
	}
	if api.ListHotelsHandler == nil {
		api.ListHotelsHandler = operations.ListHotelsHandlerFunc(func(params operations.ListHotelsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListHotels has not yet been implemented")
		})
	}
	if api.ListRoomsHandler == nil {
		api.ListRoomsHandler = operations.ListRoomsHandlerFunc(func(params operations.ListRoomsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListRooms has not yet been implemented")
		})
	}
	if api.ListUsersHandler == nil {
		api.ListUsersHandler = operations.ListUsersHandlerFunc(func(params operations.ListUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListUsers has not yet been implemented")
		})
	}
	if api.UpdateBookingHandler == nil {
		api.UpdateBookingHandler = operations.UpdateBookingHandlerFunc(func(params operations.UpdateBookingParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateBooking has not yet been implemented")
		})
	}
	if api.UpdateCarHandler == nil {
		api.UpdateCarHandler = operations.UpdateCarHandlerFunc(func(params operations.UpdateCarParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateCar has not yet been implemented")
		})
	}
	if api.UpdateHotelHandler == nil {
		api.UpdateHotelHandler = operations.UpdateHotelHandlerFunc(func(params operations.UpdateHotelParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateHotel has not yet been implemented")
		})
	}
	if api.UpdateRoomHandler == nil {
		api.UpdateRoomHandler = operations.UpdateRoomHandlerFunc(func(params operations.UpdateRoomParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateRoom has not yet been implemented")
		})
	}
	if api.UpdateUserHandler == nil {
		api.UpdateUserHandler = operations.UpdateUserHandlerFunc(func(params operations.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateUser has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

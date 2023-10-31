// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"travel-booking-portal/gen/restapi/operations"
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

	if api.AllBookingsHandler == nil {
		api.AllBookingsHandler = operations.AllBookingsHandlerFunc(func(params operations.AllBookingsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AllBookings has not yet been implemented")
		})
	}
	if api.ConfirmedBookingHandler == nil {
		api.ConfirmedBookingHandler = operations.ConfirmedBookingHandlerFunc(func(params operations.ConfirmedBookingParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ConfirmedBooking has not yet been implemented")
		})
	}
	if api.CreateOrUpdateBookingHandler == nil {
		api.CreateOrUpdateBookingHandler = operations.CreateOrUpdateBookingHandlerFunc(func(params operations.CreateOrUpdateBookingParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateOrUpdateBooking has not yet been implemented")
		})
	}
	if api.CreateOrUpdateCarHandler == nil {
		api.CreateOrUpdateCarHandler = operations.CreateOrUpdateCarHandlerFunc(func(params operations.CreateOrUpdateCarParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateOrUpdateCar has not yet been implemented")
		})
	}
	if api.CreateOrUpdateHotelHandler == nil {
		api.CreateOrUpdateHotelHandler = operations.CreateOrUpdateHotelHandlerFunc(func(params operations.CreateOrUpdateHotelParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateOrUpdateHotel has not yet been implemented")
		})
	}
	if api.CreateOrUpdateRoomHandler == nil {
		api.CreateOrUpdateRoomHandler = operations.CreateOrUpdateRoomHandlerFunc(func(params operations.CreateOrUpdateRoomParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateOrUpdateRoom has not yet been implemented")
		})
	}
	if api.CreateOrUpdateUserHandler == nil {
		api.CreateOrUpdateUserHandler = operations.CreateOrUpdateUserHandlerFunc(func(params operations.CreateOrUpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateOrUpdateUser has not yet been implemented")
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
	if api.GetBookingsHandler == nil {
		api.GetBookingsHandler = operations.GetBookingsHandlerFunc(func(params operations.GetBookingsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetBookings has not yet been implemented")
		})
	}
	if api.GetCarHandler == nil {
		api.GetCarHandler = operations.GetCarHandlerFunc(func(params operations.GetCarParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetCar has not yet been implemented")
		})
	}
	if api.GetCarsHandler == nil {
		api.GetCarsHandler = operations.GetCarsHandlerFunc(func(params operations.GetCarsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetCars has not yet been implemented")
		})
	}
	if api.GetHotelHandler == nil {
		api.GetHotelHandler = operations.GetHotelHandlerFunc(func(params operations.GetHotelParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetHotel has not yet been implemented")
		})
	}
	if api.GetHotelsHandler == nil {
		api.GetHotelsHandler = operations.GetHotelsHandlerFunc(func(params operations.GetHotelsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetHotels has not yet been implemented")
		})
	}
	if api.GetRoomHandler == nil {
		api.GetRoomHandler = operations.GetRoomHandlerFunc(func(params operations.GetRoomParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetRoom has not yet been implemented")
		})
	}
	if api.GetRoomsHandler == nil {
		api.GetRoomsHandler = operations.GetRoomsHandlerFunc(func(params operations.GetRoomsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetRooms has not yet been implemented")
		})
	}
	if api.GetServiceBookingsHandler == nil {
		api.GetServiceBookingsHandler = operations.GetServiceBookingsHandlerFunc(func(params operations.GetServiceBookingsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetServiceBookings has not yet been implemented")
		})
	}
	if api.GetUserHandler == nil {
		api.GetUserHandler = operations.GetUserHandlerFunc(func(params operations.GetUserParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetUser has not yet been implemented")
		})
	}
	if api.GetUsersHandler == nil {
		api.GetUsersHandler = operations.GetUsersHandlerFunc(func(params operations.GetUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetUsers has not yet been implemented")
		})
	}
	if api.SearchBookingHandler == nil {
		api.SearchBookingHandler = operations.SearchBookingHandlerFunc(func(params operations.SearchBookingParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.SearchBooking has not yet been implemented")
		})
	}
	if api.GetRoomsOfHotelHandler == nil {
		api.GetRoomsOfHotelHandler = operations.GetRoomsOfHotelHandlerFunc(func(params operations.GetRoomsOfHotelParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetRoomsOfHotel has not yet been implemented")
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

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"travel-booking-portal/gen/restapi/operations/bookings"
)

// NewTravelBookingPortalAPI creates a new TravelBookingPortal instance
func NewTravelBookingPortalAPI(spec *loads.Document) *TravelBookingPortalAPI {
	return &TravelBookingPortalAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		CreateBookingHandler: CreateBookingHandlerFunc(func(params CreateBookingParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateBooking has not yet been implemented")
		}),
		CreateCarHandler: CreateCarHandlerFunc(func(params CreateCarParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateCar has not yet been implemented")
		}),
		CreateHotelHandler: CreateHotelHandlerFunc(func(params CreateHotelParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateHotel has not yet been implemented")
		}),
		CreateRoomHandler: CreateRoomHandlerFunc(func(params CreateRoomParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateRoom has not yet been implemented")
		}),
		CreateUserHandler: CreateUserHandlerFunc(func(params CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateUser has not yet been implemented")
		}),
		DeleteBookingHandler: DeleteBookingHandlerFunc(func(params DeleteBookingParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteBooking has not yet been implemented")
		}),
		DeleteCarHandler: DeleteCarHandlerFunc(func(params DeleteCarParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteCar has not yet been implemented")
		}),
		DeleteHotelHandler: DeleteHotelHandlerFunc(func(params DeleteHotelParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteHotel has not yet been implemented")
		}),
		DeleteRoomHandler: DeleteRoomHandlerFunc(func(params DeleteRoomParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteRoom has not yet been implemented")
		}),
		DeleteUserHandler: DeleteUserHandlerFunc(func(params DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteUser has not yet been implemented")
		}),
		GetBookingHandler: GetBookingHandlerFunc(func(params GetBookingParams) middleware.Responder {
			return middleware.NotImplemented("operation GetBooking has not yet been implemented")
		}),
		GetCarHandler: GetCarHandlerFunc(func(params GetCarParams) middleware.Responder {
			return middleware.NotImplemented("operation GetCar has not yet been implemented")
		}),
		GetHotelHandler: GetHotelHandlerFunc(func(params GetHotelParams) middleware.Responder {
			return middleware.NotImplemented("operation GetHotel has not yet been implemented")
		}),
		GetRoomHandler: GetRoomHandlerFunc(func(params GetRoomParams) middleware.Responder {
			return middleware.NotImplemented("operation GetRoom has not yet been implemented")
		}),
		GetUserHandler: GetUserHandlerFunc(func(params GetUserParams) middleware.Responder {
			return middleware.NotImplemented("operation GetUser has not yet been implemented")
		}),
		ListBookingsHandler: ListBookingsHandlerFunc(func(params ListBookingsParams) middleware.Responder {
			return middleware.NotImplemented("operation ListBookings has not yet been implemented")
		}),
		BookingsListBookingsByDateTimeHandler: bookings.ListBookingsByDateTimeHandlerFunc(func(params bookings.ListBookingsByDateTimeParams) middleware.Responder {
			return middleware.NotImplemented("operation bookings.ListBookingsByDateTime has not yet been implemented")
		}),
		ListCarsHandler: ListCarsHandlerFunc(func(params ListCarsParams) middleware.Responder {
			return middleware.NotImplemented("operation ListCars has not yet been implemented")
		}),
		BookingsListConfirmedBookingsHandler: bookings.ListConfirmedBookingsHandlerFunc(func(params bookings.ListConfirmedBookingsParams) middleware.Responder {
			return middleware.NotImplemented("operation bookings.ListConfirmedBookings has not yet been implemented")
		}),
		ListHotelsHandler: ListHotelsHandlerFunc(func(params ListHotelsParams) middleware.Responder {
			return middleware.NotImplemented("operation ListHotels has not yet been implemented")
		}),
		ListRoomsHandler: ListRoomsHandlerFunc(func(params ListRoomsParams) middleware.Responder {
			return middleware.NotImplemented("operation ListRooms has not yet been implemented")
		}),
		ListUsersHandler: ListUsersHandlerFunc(func(params ListUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation ListUsers has not yet been implemented")
		}),
		UpdateBookingHandler: UpdateBookingHandlerFunc(func(params UpdateBookingParams) middleware.Responder {
			return middleware.NotImplemented("operation UpdateBooking has not yet been implemented")
		}),
		UpdateCarHandler: UpdateCarHandlerFunc(func(params UpdateCarParams) middleware.Responder {
			return middleware.NotImplemented("operation UpdateCar has not yet been implemented")
		}),
		UpdateHotelHandler: UpdateHotelHandlerFunc(func(params UpdateHotelParams) middleware.Responder {
			return middleware.NotImplemented("operation UpdateHotel has not yet been implemented")
		}),
		UpdateRoomHandler: UpdateRoomHandlerFunc(func(params UpdateRoomParams) middleware.Responder {
			return middleware.NotImplemented("operation UpdateRoom has not yet been implemented")
		}),
		UpdateUserHandler: UpdateUserHandlerFunc(func(params UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UpdateUser has not yet been implemented")
		}),
	}
}

/*TravelBookingPortalAPI Travel Booking Portal API (Microservices) */
type TravelBookingPortalAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// CreateBookingHandler sets the operation handler for the create booking operation
	CreateBookingHandler CreateBookingHandler
	// CreateCarHandler sets the operation handler for the create car operation
	CreateCarHandler CreateCarHandler
	// CreateHotelHandler sets the operation handler for the create hotel operation
	CreateHotelHandler CreateHotelHandler
	// CreateRoomHandler sets the operation handler for the create room operation
	CreateRoomHandler CreateRoomHandler
	// CreateUserHandler sets the operation handler for the create user operation
	CreateUserHandler CreateUserHandler
	// DeleteBookingHandler sets the operation handler for the delete booking operation
	DeleteBookingHandler DeleteBookingHandler
	// DeleteCarHandler sets the operation handler for the delete car operation
	DeleteCarHandler DeleteCarHandler
	// DeleteHotelHandler sets the operation handler for the delete hotel operation
	DeleteHotelHandler DeleteHotelHandler
	// DeleteRoomHandler sets the operation handler for the delete room operation
	DeleteRoomHandler DeleteRoomHandler
	// DeleteUserHandler sets the operation handler for the delete user operation
	DeleteUserHandler DeleteUserHandler
	// GetBookingHandler sets the operation handler for the get booking operation
	GetBookingHandler GetBookingHandler
	// GetCarHandler sets the operation handler for the get car operation
	GetCarHandler GetCarHandler
	// GetHotelHandler sets the operation handler for the get hotel operation
	GetHotelHandler GetHotelHandler
	// GetRoomHandler sets the operation handler for the get room operation
	GetRoomHandler GetRoomHandler
	// GetUserHandler sets the operation handler for the get user operation
	GetUserHandler GetUserHandler
	// ListBookingsHandler sets the operation handler for the list bookings operation
	ListBookingsHandler ListBookingsHandler
	// BookingsListBookingsByDateTimeHandler sets the operation handler for the list bookings by date time operation
	BookingsListBookingsByDateTimeHandler bookings.ListBookingsByDateTimeHandler
	// ListCarsHandler sets the operation handler for the list cars operation
	ListCarsHandler ListCarsHandler
	// BookingsListConfirmedBookingsHandler sets the operation handler for the list confirmed bookings operation
	BookingsListConfirmedBookingsHandler bookings.ListConfirmedBookingsHandler
	// ListHotelsHandler sets the operation handler for the list hotels operation
	ListHotelsHandler ListHotelsHandler
	// ListRoomsHandler sets the operation handler for the list rooms operation
	ListRoomsHandler ListRoomsHandler
	// ListUsersHandler sets the operation handler for the list users operation
	ListUsersHandler ListUsersHandler
	// UpdateBookingHandler sets the operation handler for the update booking operation
	UpdateBookingHandler UpdateBookingHandler
	// UpdateCarHandler sets the operation handler for the update car operation
	UpdateCarHandler UpdateCarHandler
	// UpdateHotelHandler sets the operation handler for the update hotel operation
	UpdateHotelHandler UpdateHotelHandler
	// UpdateRoomHandler sets the operation handler for the update room operation
	UpdateRoomHandler UpdateRoomHandler
	// UpdateUserHandler sets the operation handler for the update user operation
	UpdateUserHandler UpdateUserHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *TravelBookingPortalAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *TravelBookingPortalAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *TravelBookingPortalAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TravelBookingPortalAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TravelBookingPortalAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TravelBookingPortalAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TravelBookingPortalAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TravelBookingPortalAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TravelBookingPortalAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TravelBookingPortalAPI
func (o *TravelBookingPortalAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.CreateBookingHandler == nil {
		unregistered = append(unregistered, "CreateBookingHandler")
	}
	if o.CreateCarHandler == nil {
		unregistered = append(unregistered, "CreateCarHandler")
	}
	if o.CreateHotelHandler == nil {
		unregistered = append(unregistered, "CreateHotelHandler")
	}
	if o.CreateRoomHandler == nil {
		unregistered = append(unregistered, "CreateRoomHandler")
	}
	if o.CreateUserHandler == nil {
		unregistered = append(unregistered, "CreateUserHandler")
	}
	if o.DeleteBookingHandler == nil {
		unregistered = append(unregistered, "DeleteBookingHandler")
	}
	if o.DeleteCarHandler == nil {
		unregistered = append(unregistered, "DeleteCarHandler")
	}
	if o.DeleteHotelHandler == nil {
		unregistered = append(unregistered, "DeleteHotelHandler")
	}
	if o.DeleteRoomHandler == nil {
		unregistered = append(unregistered, "DeleteRoomHandler")
	}
	if o.DeleteUserHandler == nil {
		unregistered = append(unregistered, "DeleteUserHandler")
	}
	if o.GetBookingHandler == nil {
		unregistered = append(unregistered, "GetBookingHandler")
	}
	if o.GetCarHandler == nil {
		unregistered = append(unregistered, "GetCarHandler")
	}
	if o.GetHotelHandler == nil {
		unregistered = append(unregistered, "GetHotelHandler")
	}
	if o.GetRoomHandler == nil {
		unregistered = append(unregistered, "GetRoomHandler")
	}
	if o.GetUserHandler == nil {
		unregistered = append(unregistered, "GetUserHandler")
	}
	if o.ListBookingsHandler == nil {
		unregistered = append(unregistered, "ListBookingsHandler")
	}
	if o.BookingsListBookingsByDateTimeHandler == nil {
		unregistered = append(unregistered, "bookings.ListBookingsByDateTimeHandler")
	}
	if o.ListCarsHandler == nil {
		unregistered = append(unregistered, "ListCarsHandler")
	}
	if o.BookingsListConfirmedBookingsHandler == nil {
		unregistered = append(unregistered, "bookings.ListConfirmedBookingsHandler")
	}
	if o.ListHotelsHandler == nil {
		unregistered = append(unregistered, "ListHotelsHandler")
	}
	if o.ListRoomsHandler == nil {
		unregistered = append(unregistered, "ListRoomsHandler")
	}
	if o.ListUsersHandler == nil {
		unregistered = append(unregistered, "ListUsersHandler")
	}
	if o.UpdateBookingHandler == nil {
		unregistered = append(unregistered, "UpdateBookingHandler")
	}
	if o.UpdateCarHandler == nil {
		unregistered = append(unregistered, "UpdateCarHandler")
	}
	if o.UpdateHotelHandler == nil {
		unregistered = append(unregistered, "UpdateHotelHandler")
	}
	if o.UpdateRoomHandler == nil {
		unregistered = append(unregistered, "UpdateRoomHandler")
	}
	if o.UpdateUserHandler == nil {
		unregistered = append(unregistered, "UpdateUserHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TravelBookingPortalAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TravelBookingPortalAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *TravelBookingPortalAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *TravelBookingPortalAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *TravelBookingPortalAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TravelBookingPortalAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the travel booking portal API
func (o *TravelBookingPortalAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *TravelBookingPortalAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/booking"] = NewCreateBooking(o.context, o.CreateBookingHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/car"] = NewCreateCar(o.context, o.CreateCarHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/hotel"] = NewCreateHotel(o.context, o.CreateHotelHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/room"] = NewCreateRoom(o.context, o.CreateRoomHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user"] = NewCreateUser(o.context, o.CreateUserHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/booking/{bookingID}"] = NewDeleteBooking(o.context, o.DeleteBookingHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/car/{carID}"] = NewDeleteCar(o.context, o.DeleteCarHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/hotel/{hotelID}"] = NewDeleteHotel(o.context, o.DeleteHotelHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/room/{roomID}"] = NewDeleteRoom(o.context, o.DeleteRoomHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/user/{userID}"] = NewDeleteUser(o.context, o.DeleteUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/booking/{bookingID}"] = NewGetBooking(o.context, o.GetBookingHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/car/{carID}"] = NewGetCar(o.context, o.GetCarHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/hotel/{hotelID}"] = NewGetHotel(o.context, o.GetHotelHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/room/{roomID}"] = NewGetRoom(o.context, o.GetRoomHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/{userID}"] = NewGetUser(o.context, o.GetUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/booking"] = NewListBookings(o.context, o.ListBookingsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/bookings/{userID}/{dateTime}"] = bookings.NewListBookingsByDateTime(o.context, o.BookingsListBookingsByDateTimeHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/car"] = NewListCars(o.context, o.ListCarsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/bookings/confirmed/{userID}"] = bookings.NewListConfirmedBookings(o.context, o.BookingsListConfirmedBookingsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/hotel"] = NewListHotels(o.context, o.ListHotelsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/room"] = NewListRooms(o.context, o.ListRoomsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user"] = NewListUsers(o.context, o.ListUsersHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/booking"] = NewUpdateBooking(o.context, o.UpdateBookingHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/car"] = NewUpdateCar(o.context, o.UpdateCarHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/hotel"] = NewUpdateHotel(o.context, o.UpdateHotelHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/room"] = NewUpdateRoom(o.context, o.UpdateRoomHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/user"] = NewUpdateUser(o.context, o.UpdateUserHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TravelBookingPortalAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *TravelBookingPortalAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *TravelBookingPortalAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *TravelBookingPortalAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *TravelBookingPortalAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}

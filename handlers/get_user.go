package handlers

import (
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/models"
	"travel-booking-portal/gen/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

// NewGetUserHandler handles a request for retrieving a user
func NewGetUser(rt *runtime.Runtime) operations.GetUserHandler {
	return &getUserHandler{rt: rt}
}

type getUserHandler struct {
	rt *runtime.Runtime
}

// Handle the get user request
func (d *getUserHandler) Handle(params operations.GetUserParams) middleware.Responder {
	users, err := d.rt.Service().GetUser(int(params.UserID))

	if err != nil {

		return operations.NewGetUserInternalServerError()
	}

	if len(users) > 0 {
		// Convert your application's User type to the Swagger-generated User type
		swaggerUser := &models.User{
			Name:  users[0].Name, // Map the fields accordingly
			Email: users[0].Email,
		}

		return operations.NewGetUserOK().WithPayload(swaggerUser)
	}

	// Handle the case when no user is found
	return operations.NewGetUserNotFound()
}

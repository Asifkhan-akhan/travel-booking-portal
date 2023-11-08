package handlers

import (
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/restapi/operations"
	"travel-booking-portal/models"

	"github.com/go-openapi/runtime/middleware"
)

// NewCreateUser handles requests for creating a user
func NewCreateUser(rt *runtime.Runtime) operations.CreateUserHandler {
	return &createUserHandler{rt: rt}
}

type createUserHandler struct {
	rt *runtime.Runtime
}

// Handle the create user request
func (d *createUserHandler) Handle(params operations.CreateUserParams) middleware.Responder {
	id, err := d.rt.Service().CreateUser(&models.User{
		Name:  params.User.Name,
		Email: params.User.Email,
	})

	if err != nil {
		log().Errorf("failed to create user: %s", err)
		if err.Error() == "Username and email are required" {
			return operations.NewCreateUserBadRequest()

		}
	}

	params.User.ID = int64(id)
	return operations.NewCreateUserCreated().WithPayload(params.User)
}

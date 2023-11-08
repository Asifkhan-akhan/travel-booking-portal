package handlers

import (
	runtime "travel-booking-portal"
	"travel-booking-portal/gen/models"
	"travel-booking-portal/gen/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

// NewListUserHandler handles a request for retrieving a user
func NewListUser(rt *runtime.Runtime) operations.ListUsersHandler {
	return &listUserHandler{rt: rt}
}

type listUserHandler struct {
	rt *runtime.Runtime
}

func (d *listUserHandler) Handle(params operations.ListUsersParams) middleware.Responder {
	users, err := d.rt.Service().ListUsers()
	if err != nil {
		return operations.NewListUsersInternalServerError()
	}

	var swaggerUsers []*models.User

	for _, user := range users {
		swaggerUsers = append(swaggerUsers, &models.User{
			ID:    int64(user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	if len(swaggerUsers) > 0 {
		return operations.NewListUsersOK().WithPayload(swaggerUsers)
	}

	return operations.NewListUsersNotFound()
}

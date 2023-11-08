package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "travel-booking-portal"

	"travel-booking-portal/gen/restapi/operations"
)

func NewDeleteUser(rt *runtime.Runtime) operations.DeleteUserHandler {
	return &deleteUser{rt: rt}
}

type deleteUser struct {
	rt *runtime.Runtime
}

// Handle the delete entry request
func (d *deleteUser) Handle(params operations.DeleteUserParams) middleware.Responder {
	if err := d.rt.Service().DeleteUser(int(params.UserID)); err != nil {

		return operations.NewDeleteUserInternalServerError()
	}

	return operations.NewDeleteUserNoContent()
}

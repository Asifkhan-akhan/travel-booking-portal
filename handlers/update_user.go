package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "travel-booking-portal"
	"travel-booking-portal/gen/restapi/operations"
	"travel-booking-portal/models"
)

func NewUpdateUser(rt *runtime.Runtime) operations.UpdateUserHandler {
	return &updateuser{rt: rt}
}

type updateuser struct {
	rt *runtime.Runtime
}

func (d *updateuser) Handle(params operations.UpdateUserParams) middleware.Responder {
	up_user := &models.User{
		ID:    int(params.UserID),
		Name:  params.User.Name,
		Email: params.User.Email,
	}
	if _, err := d.rt.Service().UpdateUser(up_user); err != nil {
		if err != nil {
			log().Errorf("failed to update  user: %s", err)
			if err.Error() == "Username and email are required" {
				return operations.NewUpdateUserBadRequest()

			}
		}
		return operations.NewUpdateUserInternalServerError()
	}

	return operations.NewUpdateUserOK()
}

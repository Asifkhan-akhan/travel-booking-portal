package runtime

import (
	"travel-booking-portal/db/mongo"
	"travel-booking-portal/service"
)

// Runtime initializes values for entry point to our application
type Runtime struct {
	svc *service.Service
}

// NewRuntime creates a new database service
func NewRuntime() (*Runtime, error) {
	store, err := mongo.NewClient()
	if err != nil {
		return nil, err
	}
	return &Runtime{svc: service.NewService(store)}, err
}

// Service returns pointer to our service variable
func (r Runtime) Service() *service.Service {
	return r.svc
}

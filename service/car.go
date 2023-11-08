package service

import "travel-booking-portal/models"

func (s *Service) CreateCarRental(car *models.CarRental) (int, error) {
	carid, err := s.db.CreateOrUpdateCarRental(car)
	if err != nil {
		return 0, err
	}
	return carid, nil
}

func (s *Service) UpdateCarRental(car *models.CarRental) (int, error) {
	carid, err := s.db.CreateOrUpdateCarRental(car)
	if err != nil {
		return 0, err
	}
	return carid, nil
}
func (s *Service) DeleteCarRental(carID int) error {
	_, err := s.GetCar(carID)
	if err != nil {
		return err
	}
	return s.db.DeleteCarRental(carID)
}

// ListCarByID retrieves a car by its ID.
func (s *Service) GetCar(carID int) ([]*models.CarRental, error) {
	searchCriteria := map[string]interface{}{"_id": carID}

	car, err := s.db.ListCarRental(searchCriteria)
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (s *Service) ListCars() ([]*models.CarRental, error) {

	cars, err := s.db.ListCarRental(nil)
	if err != nil {
		return nil, err
	}

	return cars, nil
}

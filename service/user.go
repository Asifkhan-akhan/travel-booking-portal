package service

import (
	"errors"
	"travel-booking-portal/models"
)

func (s *Service) CreateUser(user *models.User) (int, error) {
	if user.Email == "" || user.Name == "nil" {
		return 0, errors.New("Username and email are required")

	}
	userid, err := s.db.CreateOrUpdateUser(user)
	if err != nil {
		return 0, err
	}
	return userid, nil
}

func (s *Service) UpdateUser(user *models.User) (int, error) {
	if user.Email == "" || user.Name == "nil" {
		return 0, errors.New("Username and email are required")

	}
	userid, err := s.db.CreateOrUpdateUser(user)
	if err != nil {
		return 0, err
	}
	return userid, nil
}

func (s *Service) GetUser(UserID int) (user []*models.User, err error) {
	user, err = s.db.ListUser(map[string]interface{}{"_id": UserID})
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *Service) ListUsers() (user []*models.User, err error) {
	user, err = s.db.ListUser(nil)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *Service) DeleteUser(userID int) error {
	err := s.db.DeleteUser(userID)
	return err
}

package service

import "travel-booking-portal/models"

func (s *Service) CreateHotel(hotel *models.Hotel) (int, error) {
	hotelid, err := s.db.CreateOrUpdateHotel(hotel)
	if err != nil {
		return 0, err
	}
	return hotelid, nil
}
func (s *Service) UpdateHotel(hotel *models.Hotel) (int, error) {
	hotelid, err := s.db.CreateOrUpdateHotel(hotel)
	if err != nil {
		return 0, err
	}
	return hotelid, nil
}
func (s *Service) GetHotel(hotelID int) ([]*models.Hotel, error) {
	hotel, err := s.db.ListHotel(map[string]interface{}{"_id": hotelID})
	if err != nil {
		return nil, err
	}
	return hotel, nil
}
func (s *Service) ListHotels() ([]*models.Hotel, error) {
	hotel, err := s.db.ListHotel(nil)
	if err != nil {
		return nil, err
	}
	return hotel, nil
}
func (s *Service) DeleteHotel(hotelID int) error {
	err := s.db.DeleteHotel(hotelID)
	if err != nil {
		return err
	}
	return s.db.DeleteHotel(hotelID)
}

func (s *Service) CreateRoom(room *models.Room) (int, error) {
	roomid, err := s.db.CreateOrUpdateRoom(room)
	if err != nil {
		return 0, err
	}
	return roomid, nil
}
func (s *Service) UpdateRoom(room *models.Room) (int, error) {
	roomid, err := s.db.CreateOrUpdateRoom(room)
	if err != nil {
		return 0, err
	}
	return roomid, nil
}
func (s *Service) GetRoom(roomID int) ([]*models.Room, error) {
	room, err := s.db.ListRoom(map[string]interface{}{"_id": roomID})
	if err != nil {
		return nil, err
	}
	return room, nil
}
func (s *Service) ListRooms() ([]*models.Room, error) {
	room, err := s.db.ListRoom(nil)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (s *Service) DeleteRoom(roomID int) error {
	err := s.db.DeleteHotel(roomID)
	if err != nil {
		return err
	}
	return s.db.DeleteRoom(roomID)
}

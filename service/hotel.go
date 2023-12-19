package service

import (
	"strconv"
	"travel-booking-portal/models"

	"github.com/pkg/errors"
)

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
	hotel, err := s.GetHotel(room.HotelID)
	if err != nil {
		return 0, err
	}

	if len(hotel) > 0 {
		//checks if the room with the same room number exists in this hotel
		oldRoom, err := s.db.ListRoom(map[string]interface{}{"number": room.Number, "hotel_id": hotel[0].ID})
		if err != nil {
			return 0, err
		}
		hotel := hotel[0]
		if len(oldRoom) == 0 {
			if roomNumber, err := strconv.Atoi(room.Number); err != nil {
				return 0, errors.Wrap(err, "failed to convert room number to integer")
			} else if roomNumber > hotel.Rooms {
				return 0, errors.New("Number of room exceeds the hotel's capacity")
			}

			roomID, err := s.db.CreateOrUpdateRoom(room)
			if err != nil {
				return 0, err
			}
			return roomID, nil
		}
		return 0, errors.New("Room with this number already exists")
	}

	return 0, errors.New("Failed to fetch the hotel")
}

func (s *Service) UpdateRoom(room *models.Room) (int, error) {
	hotel, _ := s.GetHotel(room.HotelID)
	if len(hotel) > 0 {
		roomid, err := s.db.CreateOrUpdateRoom(room)
		if err != nil {
			return 0, err
		}
		return roomid, nil
	}
	return 0, errors.New("Failed to fetch the hotel")
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

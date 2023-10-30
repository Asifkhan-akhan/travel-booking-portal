package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"travel-booking-portal/config"
	"travel-booking-portal/db"

	"travel-booking-portal/models"
)

const (
	carCollection   = "Car"
	hotelCollection = "Hotel"
	roomCollection  = "Room"
	userCollection  = "User"
	bokCollection   = "Booking"
)

type client struct {
	conn *mongo.Client
}

// NewClient initializes a mysql database connection
func NewClient() (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	log().Infof("initializing mongodb: %s", uri)
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {

		return nil, errors.Wrap(err, "failed to connect to db")
	}

	return &client{conn: cli}, nil
}

func (c *client) CreateOrUpdateBooking(booking *models.Booking) (int, error) {
	if booking.ID == 0 {

		return 0, errors.New("ID is not empty")
	}
	if booking.ServiceType != "Hotel" && booking.ServiceType != "CarRental" {

		return 0, errors.New("Invalid service type passed")
	}

	overlappingFilter := bson.D{
		{Key: "service_id", Value: booking.ServiceID},
		{Key: "from_date", Value: bson.D{{Key: "$lt", Value: booking.ToDate}}},
		{Key: "to_date", Value: bson.D{{Key: "$gt", Value: booking.FromDate}}}}
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(bokCollection)
	filter := bson.D{{Key: "_id", Value: booking.ID}}
	update := bson.D{{Key: "$set", Value: booking}}
	// overlapCount, _ := collection.CountDocuments(context.TODO(), overlappingFilter)
	result := collection.FindOne(context.TODO(), overlappingFilter)
	//if no booking is find , create new one
	if result.Err() == mongo.ErrNoDocuments {
		_, err := collection.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))
		if err != nil {

			return 0, errors.Wrap(err, "failed to update or create booking")
		}
	} else {
		//update
		if result.Err() == nil {
			var dbbooking models.Booking
			if err := result.Decode(&dbbooking); err != nil {
				errors.New("Error decoding result:" + err.Error())
			}
			_, err := collection.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(false))
			if err != nil {

				return 0, errors.Wrap(err, "failed to create booking")
			}
		}
	}

	return booking.ID, nil
}

func (c *client) CreateOrUpdateCarRental(car *models.CarRental) (int, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(carCollection)
	filter := bson.D{{Key: "_id", Value: car.ID}}

	update := bson.D{{Key: "$set", Value: car}}

	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {

		return 0, errors.Wrap(err, "failed to update or create car rental")
	}

	return car.ID, nil
}

func (c *client) CreateOrUpdateHotel(hotel *models.Hotel) (int, error) {
	if hotel.ID == 0 {

		return 0, errors.New("ID can't be empty")
	}

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(hotelCollection)

	update := bson.D{{"$set", hotel}}

	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.TODO(), bson.D{{"_id", hotel.ID}}, update, opts)

	if err != nil {

		return 0, errors.Wrap(err, "failed to update or create hotel")
	}

	return hotel.ID, nil
}

func (c *client) CreateOrUpdateRoom(room *models.Room) (int, error) {
	if room.HotelID == 0 {

		return 0, errors.New("hotel id can't be empty")
	}

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(roomCollection)

	filter := bson.D{{"_id", room.ID}}

	update := bson.D{{"$set", room}}

	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {

		return 0, errors.Wrap(err, "failed to update or create room")
	}

	return room.ID, nil
}

func (c *client) CreateOrUpdateUser(user *models.User) (int, error) {
	if user.Name == "" || user.Email == "" {

		return 0, errors.New("Name and Email can't be empty")
	}

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(userCollection)

	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(context.TODO(), bson.D{{"_id", user.ID}}, bson.D{{"$set", user}}, opts)

	if err != nil {

		return 0, errors.Wrap(err, "failed to update or create user")
	}

	return user.ID, nil
}

// DeleteBooking implements db.DataStore.
func (c *client) DeleteBooking(bookingID int) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(bokCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": bookingID}); err != nil {

		return errors.Wrap(err, "failed to delete Booking")
	}

	return nil
}

// DeleteCarRental implements db.DataStore.
func (c *client) DeleteCarRental(carID int) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(carCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": carID}); err != nil {

		return errors.Wrap(err, "failed to delete Car")
	}

	return nil
}

// DeleteHotel implements db.DataStore.
func (c *client) DeleteHotel(hotelID int) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(hotelCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": hotelID}); err != nil {

		return errors.Wrap(err, "failed to delete hotel")
	}

	return nil
}

// DeleteRoom implements db.DataStore.
func (c *client) DeleteRoom(roomID int) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(roomCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": roomID}); err != nil {

		return errors.Wrap(err, "failed to delete room")
	}

	return nil
}

// DeleteUser implements db.DataStore.
func (c *client) DeleteUser(userID int) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(userCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": userID}); err != nil {

		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}

// GetBooking implements db.DataStore.
func (c *client) GetBooking(bookingID int) ([]*models.Booking, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(bokCollection)
	var bookings []*models.Booking
	var cursor *mongo.Cursor
	var err error

	if bookingID == 0 {
		cursor, err = collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{{"_id", 1}}))
	} else {
		// If bookingID is not 0, we're retrieving a specific booking
		cursor, err = collection.Find(context.TODO(), bson.M{"_id": bookingID})
	}

	if err != nil {

		return nil, err
	}

	if err := cursor.All(context.TODO(), &bookings); err != nil {

		return nil, err
	}

	return bookings, nil
}

// GetCarRental implements db.DataStore.
func (c *client) GetCarRental(carID int) ([]*models.CarRental, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(carCollection)
	var cars []*models.CarRental
	var cursor *mongo.Cursor
	var err error

	if carID == 0 {
		// If bookingID is 0, we're retrieving all cars
		cursor, err = collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{{"_id", 1}}))
		fmt.Println("running id 0")
	} else {
		// If bookingID is not 0, we're retrieving a specific car
		cursor, err = collection.Find(context.TODO(), bson.M{"_id": carID})
	}

	if err != nil {

		return nil, err

	}

	if err := cursor.All(context.TODO(), &cars); err != nil {
		fmt.Print("theser is and error\n", err)

		return nil, err
	}

	return cars, nil
}

// GetHotel implements db.DataStore.
func (c *client) GetHotel(hotelID int) ([]*models.Hotel, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(hotelCollection)
	var hotel []*models.Hotel
	var cursor *mongo.Cursor
	var err error

	if hotelID == 0 {
		// If bookingID is 0, we're retrieving all cars
		cursor, err = collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{{"_id", 1}}))
	} else {
		// If bookingID is not 0, we're retrieving a specific car
		cursor, err = collection.Find(context.TODO(), bson.M{"_id": hotelID})
	}

	if err != nil {

		return nil, err
	}

	if err := cursor.All(context.TODO(), &hotel); err != nil {

		return nil, err
	}

	return hotel, nil
}

// GetRoom implements db.DataStore.
func (c *client) GetRoom(roomID int) ([]*models.Room, error) {

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(roomCollection)
	var room []*models.Room
	var cursor *mongo.Cursor
	var err error

	if roomID == 0 {
		cursor, err = collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{{"_id", 1}}))
	} else {
		cursor, err = collection.Find(context.TODO(), bson.M{"_id": roomID})
	}

	if err != nil {

		return nil, err
	}

	if err := cursor.All(context.TODO(), &room); err != nil {

		return nil, err
	}

	return room, nil
}

// GetUser implements db.DataStore.
func (c *client) GetUser(userID int) ([]*models.User, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(userCollection)
	var user []*models.User
	var cursor *mongo.Cursor
	var err error

	if userID == 0 {
		// If bookingID is 0, we're retrieving all cars
		cursor, err = collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{{Key: "_id", Value: 1}}))
	} else {
		// If bookingID is not 0, we're retrieving a specific car
		cursor, err = collection.Find(context.TODO(), bson.M{"_id": userID})
	}

	if err != nil {

		return nil, err
	}

	if err := cursor.All(context.TODO(), &user); err != nil {

		return nil, err
	}

	return user, nil
}
func (c *client) ConfirmedBooking(userid int) ([]*models.Booking, error) {
	// Define a filter to match bookings with the specified user ID and date.
	filter := bson.D{
		{Key: "user_id", Value: userid},
		{Key: "confirmed", Value: true},
	}

	// Create a slice to store the results.
	var bookings []*models.Booking

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(bokCollection)

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {

		return nil, errors.Wrap(err, "failed to search for bookings")
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var booking models.Booking
		if err := cursor.Decode(&booking); err != nil {

			return nil, errors.Wrap(err, "failed to decode booking")
		}
		bookings = append(bookings, &booking)
	}

	if err := cursor.Err(); err != nil {

		return nil, errors.Wrap(err, "cursor error")
	}

	return bookings, nil
}
func (c *client) SearchBooking(userid int, dateString string) ([]*models.Booking, error) {
	dateTime, err := time.Parse("2006, 01, 02", dateString)
	if err != nil {

		return nil, errors.Wrap(err, "failed to parse date string")
	}

	startOfDay := time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 0, 0, 0, 0, time.UTC)

	filter := bson.D{
		{Key: "user_id", Value: userid},
		{Key: "from_date", Value: startOfDay},
	}

	var bookings []*models.Booking

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(bokCollection)

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {

		return nil, errors.Wrap(err, "failed to search for bookings")
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var booking models.Booking
		if err := cursor.Decode(&booking); err != nil {

			return nil, errors.Wrap(err, "failed to decode booking")
		}
		bookings = append(bookings, &booking)
	}

	if err := cursor.Err(); err != nil {

		return nil, errors.Wrap(err, "cursor error")
	}

	return bookings, nil
}

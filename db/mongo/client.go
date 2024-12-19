package mongo

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"travel-booking-portal/db"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"travel-booking-portal/config"
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
var (
	once     sync.Once
	instance db.DataStore
)

func NewClient() (db.DataStore, error) {
	once.Do(func() {
		uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
		log().Infof("Initializing MongoDB: %s", uri)

		cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			log().Errorf("Failed to connect to MongoDB: %v", err)
			return
		}

		instance = &client{conn: cli}
	})

	return instance, nil
}

func generateUniqueID() int {
	u := uuid.New()
	// Convert UUID to a big.Int
	idInt := new(big.Int)
	idInt.SetString(u.String(), 16)

	// Convert the big.Int to an integer
	return int(idInt.Int64())
}

func (c *client) CreateOrUpdateBooking(booking *models.Booking) (int, error) {

	if booking.ServiceType != "Hotel" && booking.ServiceType != "CarRental" {

		return 0, errors.New("Invalid service type passed")
	}
	overlappingFilter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.D{{Key: "service_id", Value: booking.ServiceID}},
				bson.D{{Key: "from_date", Value: bson.D{{Key: "$lt", Value: booking.ToDate}}}},
				bson.D{{Key: "to_date", Value: bson.D{{Key: "$gt", Value: booking.FromDate}}}},
			},
		},
	}

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(bokCollection)
	filter := bson.D{{Key: "_id", Value: booking.ID}}
	update := bson.D{{Key: "$set", Value: booking}}
	result := collection.FindOne(context.TODO(), overlappingFilter)
	//if no booking is find , create new one
	if result.Err() == mongo.ErrNoDocuments {
		//create new
		booking.ID = generateUniqueID()
		_, err := collection.InsertOne(context.TODO(), booking)
		if err != nil {

			return 0, errors.Wrap(err, "failed to create booking")
		}
		return booking.ID, nil

	} else {
		//update
		if result.Err() == nil {
			var dbbooking models.Booking

			if err := result.Decode(&dbbooking); err != nil {
				return 0, errors.Wrap(err, "Error decoding result:"+err.Error())

			}

			_, err := collection.UpdateOne(context.TODO(), filter, update, options.Update())
			if err != nil {

				return 0, errors.Wrap(err, "failed to update booking")
			}
			return booking.ID, nil

		}
	}

	return 0, nil
}

func (c *client) CreateOrUpdateCarRental(car *models.CarRental) (int, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(carCollection)

	if car.ID != 0 {

		filter := bson.D{{Key: "_id", Value: car.ID}}

		opts := options.Update().SetUpsert(true)

		_, err := collection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: car}}, opts)
		if err != nil {

			return 0, errors.Wrap(err, "failed to update car rental")

		}
	} else {
		// Generate a unique ID for the new car rental.
		newID := generateUniqueID()
		car.ID = newID

		_, err := collection.InsertOne(context.TODO(), car)
		if err != nil {
			return 0, errors.Wrap(err, "failed to create car rental")
		}
	}
	return car.ID, nil

}

func (c *client) CreateOrUpdateHotel(hotel *models.Hotel) (int, error) {
	if hotel.ID != 0 {

		collection := c.conn.Database(viper.GetString(config.DbName)).Collection(hotelCollection)

		//update := bson.D{{"$set", hotel}}
		update := bson.D{{Key: "$set", Value: hotel}}

		opts := options.Update().SetUpsert(true)

		_, err := collection.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: hotel.ID}}, update, opts)

		if err != nil {
			return 0, errors.Wrap(err, "failed to update hotel")
		}

		return hotel.ID, nil
	}

	// Generate a unique ID for the new hotel.
	newID := generateUniqueID()
	hotel.ID = newID

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(hotelCollection)

	_, err := collection.InsertOne(context.TODO(), hotel)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create hotel")
	}

	return newID, nil
}

func (c *client) CreateOrUpdateRoom(room *models.Room) (int, error) {
	if room.ID != 0 {
		// ID is not empty; this is an update, not a creation.
		collection := c.conn.Database(viper.GetString(config.DbName)).Collection(roomCollection)
		update := bson.D{{Key: "$set", Value: room}}
		opts := options.Update().SetUpsert(true)

		_, err := collection.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: room.ID}}, update, opts)
		if err != nil {
			return 0, errors.Wrap(err, "failed to update room")
		}

		return room.ID, nil
	}

	// Generate a unique ID for the new room.
	newID := generateUniqueID()
	room.ID = newID

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(roomCollection)

	_, err := collection.InsertOne(context.TODO(), room)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create room")
	}

	return newID, nil
}

func (c *client) CreateOrUpdateUser(user *models.User) (int, error) {
	if user.ID != 0 {
		// ID is not empty; this is an update, not a creation.
		collection := c.conn.Database(viper.GetString(config.DbName)).Collection(userCollection)
		opts := options.Update().SetUpsert(true)

		_, err := collection.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: user.ID}}, bson.D{{Key: "$set", Value: user}}, opts)
		if err != nil {
			return 0, errors.Wrap(err, "failed to update user")
		}

		return user.ID, nil
	}

	// Generate a unique ID for the new user.
	newID := generateUniqueID()
	user.ID = newID

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(userCollection)

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create user")
	}

	return newID, nil
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

// ListBooking retrieves bookings based on the provided search parameters.
func (c *client) ListBooking(searchCriteria map[string]interface{}) ([]*models.Booking, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(bokCollection)
	var bookings []*models.Booking
	var cursor *mongo.Cursor
	var err error

	// Create a filter based on the search parameters
	filter := bson.M{}

	for key, value := range searchCriteria {
		filter[key] = value
	}

	cursor, err = collection.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{Key: "_id", Value: 1}}))

	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &bookings); err != nil {
		return nil, err
	}

	return bookings, nil
}

// ListCarRental implements db.DataStore.
func (c *client) ListCarRental(searchCriteria map[string]interface{}) ([]*models.CarRental, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(carCollection)
	var cars []*models.CarRental
	var cursor *mongo.Cursor
	var err error

	filter := bson.M{}

	// Populate the filter based on the provided search criteria.

	for key, value := range searchCriteria {
		filter[key] = value
	}

	cursor, err = collection.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{Key: "_id", Value: 1}}))
	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &cars); err != nil {
		return nil, err
	}

	return cars, nil
}

// ListHotel implements db.DataStore.
func (c *client) ListHotel(searchCriteria map[string]interface{}) ([]*models.Hotel, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(hotelCollection)
	var hotels []*models.Hotel
	var cursor *mongo.Cursor
	var err error

	filter := bson.M{}

	for key, value := range searchCriteria {
		filter[key] = value
	}

	cursor, err = collection.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{Key: "_id", Value: 1}}))

	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &hotels); err != nil {
		return nil, err
	}

	return hotels, nil
}

// ListRoom implements db.DataStore.
func (c *client) ListRoom(searchCriteria map[string]interface{}) ([]*models.Room, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(roomCollection)
	var rooms []*models.Room
	var cursor *mongo.Cursor
	var err error

	filter := bson.M{}

	for key, value := range searchCriteria {
		filter[key] = value
	}

	cursor, err = collection.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{Key: "_id", Value: 1}}))
	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &rooms); err != nil {
		return nil, err
	}

	return rooms, nil
}

// ListUser implements db.DataStore.
func (c *client) ListUser(searchCriteria map[string]interface{}) ([]*models.User, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(userCollection)
	var users []*models.User
	var cursor *mongo.Cursor
	var err error

	filter := bson.M{}

	for key, value := range searchCriteria {
		filter[key] = value
	}

	cursor, err = collection.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{Key: "_id", Value: 1}}))
	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

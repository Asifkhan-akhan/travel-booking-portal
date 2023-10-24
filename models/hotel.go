package models

import "github.com/fatih/structs"

type Hotel struct {
	ID          int    `json:"id" db:"id" structs:"id" bson:"_id"`
	Name        string `json:"name" db:"name" structs:"name" bson:"name"`
	Location    string `json:"location" db:"location" structs:"location" bson:"location"`
	Description string `json:"description" db:"description" structs:"description" bson:"description"`
	Rooms       int    `json:"rooms" db:"rooms" structs:"rooms" bson:"rooms"`
}

type Room struct {
	ID      int    `json:"id" db:"id" structs:"id" bson:"_id"`
	Number  string `json:"number" db:"number" structs:"number" bson:"number"`
	HotelID int    `json:"hotel_id" db:"hotel_id" structs:"hotel_id" bson:"hotel_id"`
}

// Map converts structs to a map representation
func (s *Hotel) Map() map[string]interface{} {
	return structs.Map(s)
}

// Names returns the field names of Hotel model
func (s *Hotel) Names() []string {
	fields := structs.Fields(s)
	names := make([]string, len(fields))

	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}

// Map converts structs to a map representation
func (s *Room) Map() map[string]interface{} {
	return structs.Map(s)
}

// Names returns the field names of Room model
func (s *Room) Names() []string {
	fields := structs.Fields(s)
	names := make([]string, len(fields))

	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}

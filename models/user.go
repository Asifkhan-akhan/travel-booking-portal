package models

import "github.com/fatih/structs"

type User struct {
	ID    int    `json:"id" db:"id" structs:"id" bson:"_id"`
	Name  string `json:"name" db:"name" structs:"name" bson:"name"`
	Email string `json:"email" db:"email" structs:"email" bson:"email"`
}

// Map converts structs to a map representation
func (s *User) Map() map[string]interface{} {
	return structs.Map(s)
}

// Names returns the field names of User model
func (s *User) Names() []string {
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

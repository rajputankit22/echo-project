package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents the structure of a user document in MongoDB
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // MongoDB auto-generated ID
	Name      string             `bson:"name" json:"name" validate:"required"`
	Email     string             `bson:"email" json:"email" validate:"required,email"`
	Password  string             `bson:"password" json:"password" validate:"required,min=6"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

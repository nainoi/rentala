package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is user data storage
type User struct {
	UserID     primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Email      string             `json:"email" bson:"email"`
	Password   string             `json:"-" bson:"password"`
	Provider   string             `json:"-" bson:"provider"`
	ProviderID string             `json:"-" bson:"provider_id"`
	FirstName  string             `json:"firstname" bson:"firstname"`
	LastName   string             `json:"lastname" bson:"lastname"`
	Role       string             `json:"role" bson:"role"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}

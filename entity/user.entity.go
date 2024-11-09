package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	UserName  string             `bson:"user_name" json:"user_name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

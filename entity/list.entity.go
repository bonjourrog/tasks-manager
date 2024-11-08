package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type List struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Color     string             `bson:"color" json:"color"`
	Tasks     []Task             `bson:"tasks" json:"tasks"`
	UserID    string             `bson:"user_id" json:"user_id"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

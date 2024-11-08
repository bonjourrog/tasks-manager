package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Done        bool               `bson:"done" json:"done"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

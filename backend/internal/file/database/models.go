package database

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     int
	Username   string
	FileName   string
	FilePath   string
	SpecialURL uuid.UUID
	ExpiryDate string
	Password   string `bson:",omitempty"`
}

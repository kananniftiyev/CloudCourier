package database

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type File struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     int
	Username   string
	FileName   string
	FilePath   string
	SpecialURL uuid.UUID
	ExpiryDate time.Time
	Password   string `bson:",omitempty"`
}

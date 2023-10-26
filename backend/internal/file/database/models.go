package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type File struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   int
	Username string
	FileName string
	FilePath string
	Password string `bson:",omitempty"`
}

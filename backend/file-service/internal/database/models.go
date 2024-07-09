package database

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID         int                `json:"user_id"`
	Username       string             `json:"username"`
	FileName       string             `json:"file_name"`
	FilePath       string             `json:"file_path"`
	SpecialURL     uuid.UUID          `json:"special_url"`
	ExpiryDate     string             `json:"expiry_date"`
	Password       string             `bson:",omitempty" json:"password"`
	Title          string             `json:"title"`
	TotalDownloads int                `json:"totalDownloads"`
	UploadDate     string             `json:"upload_date"`
}

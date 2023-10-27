package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FileRepository struct {
	collection *mongo.Collection
}

func NewFileRepository(db *mongo.Database) *FileRepository {
	return &FileRepository{
		collection: db.Collection("Files"),
	}
}

func (r *FileRepository) Create(ctx context.Context, file *File) error {
	_, err := r.collection.InsertOne(ctx, file)
	return err
}

func (r *FileRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*File, error) {
	var file File
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&file)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (r *FileRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

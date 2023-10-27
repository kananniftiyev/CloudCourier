package database

import (
	"backend/internal/auth/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	db := ConnectDatabase()
	db.AutoMigrate(&models.User{})
	db.Commit()
}

func ConnectDatabase() *gorm.DB {
	dsn := "user=postgres password=kanan123 dbname=CloudShareX host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to Database")
	}
	return db
}

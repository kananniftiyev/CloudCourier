package database

import (
	"backend/auth-service/internal/database/models"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func init() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	envFile := filepath.Join(dir, ".env")
	err = godotenv.Load(envFile)
	if err != nil {
		log.Fatal(err)
	}
	db := ConnectDatabase()
	db.AutoMigrate(&models.User{})
	db.Commit()
}

func ConnectDatabase() *gorm.DB {
	var port = os.Getenv("DB_PORT")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic("Failed to convert port to an integer")
	}
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), portInt, os.Getenv("DB_NAME"))
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to Database")
	}
	return db
}

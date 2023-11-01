package database

import (
	"backend/internal/auth/database/models"
	"database/sql"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"os"
	"strconv"
)

var db *sql.DB
var server = os.Getenv("DB_HOST")
var port = os.Getenv("DB_PORT")
var user = os.Getenv("DB_USER")
var password = os.Getenv("DB_PASSWORD")
var database = os.Getenv("DB_NAME")

func init() {
	db := ConnectDatabase()
	db.AutoMigrate(&models.User{})
	db.Commit()
}

func ConnectDatabase() *gorm.DB {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic("Failed to convert port to an integer")
	}
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", user, password, server, portInt, database)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to Database")
	}
	return db
}

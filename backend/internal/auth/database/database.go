package database

import (
	"backend/internal/auth/database/models"
	"database/sql"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db *sql.DB
var server = "kananniftiyev.database.windows.net"
var port = 1433
var user = "post"
var password = "Kanan19102004k"
var database = "Users"

func init() {
	db := ConnectDatabase()
	db.AutoMigrate(&models.User{})
	db.Commit()
}

func ConnectDatabase() *gorm.DB {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", user, password, server, port, database)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to Database")
	}
	return db
}

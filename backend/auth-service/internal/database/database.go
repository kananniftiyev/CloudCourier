package database

import (
	"backend/auth-service/internal/database/models"
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitUserScheme() {
	db := ConnectDatabase()
	db.AutoMigrate(&models.User{})
	db.Commit()
}

var databaseInstance *gorm.DB
var databaseLock sync.Mutex

// ConnectDatabase establishes a connection to the database and returns a pointer to the gorm.DB instance.
// If a connection has already been established, it returns the existing instance.
func ConnectDatabase() *gorm.DB {
	if databaseInstance == nil {
		databaseLock.Lock()
		defer databaseLock.Unlock()
		if os.Getenv("DB_HOST") == "" {
			log.Fatal("Make sure enter all Envoriments variables.")
		}
		var port = os.Getenv("DB_PORT")
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), port)
		databaseInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed to Connect to Database")
		}
		return databaseInstance
	}

	return databaseInstance
}

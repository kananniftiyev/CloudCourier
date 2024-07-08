package database

import (
	"backend/auth-service/internal/database/models"
	"fmt"
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

func ConnectDatabase() *gorm.DB {
	if databaseInstance == nil {
		databaseLock.Lock()
		defer databaseLock.Unlock()
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

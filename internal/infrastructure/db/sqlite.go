package db

import (
	"kaung-htet-hein-dev/finance-tracker-go/internal/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.User{})

	return db
}

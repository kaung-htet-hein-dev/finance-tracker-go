package db

import (
	"kaung-htet-hein-dev/finance-tracker-go/internal/domain"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data/finance.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Transaction{})

	return db
}

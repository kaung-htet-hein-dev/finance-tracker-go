package database

import (
	"fmt"
	"log"

	"finance-tracker-go/internal/config"
	"finance-tracker-go/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(cfg *config.Config) error {
	var err error
	var dsn string

	switch cfg.Database.Type {
	case "sqlite":
		dsn = cfg.Database.Name
		DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			cfg.Database.Host,
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Name,
			cfg.Database.Port,
			cfg.Database.SSLMode)
		// Note: For PostgreSQL, you would need to import the driver:
		// "gorm.io/driver/postgres" and use postgres.Open(dsn)
		// But we'll stick with SQLite for simplicity
		return fmt.Errorf("postgres support not implemented in this example")
	default:
		return fmt.Errorf("unsupported database type: %s", cfg.Database.Type)
	}

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully")
	return nil
}

func Migrate() error {
	if DB == nil {
		return fmt.Errorf("database connection not established")
	}

	err := DB.AutoMigrate(
		&models.User{},
		&models.Transaction{},
	)

	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
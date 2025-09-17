package pkg

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

func HandleGormError(err error, entity string) error {

	if err == nil {
		return nil
	}

	errMsg := err.Error()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return errors.New(entity + " not found")

	case errors.Is(err, gorm.ErrDuplicatedKey):
		return errors.New(entity + " already exists")

	// Handle sqlite3 UNIQUE constraint error
	case errMsg == "UNIQUE constraint failed: users.email":
		return errors.New(entity + " already exists")

	// Handle generic UNIQUE constraint errors for SQLite
	case strings.Contains(errMsg, "UNIQUE constraint failed"):
		return errors.New(entity + " already exists")

	default:
		return errors.New("database error: " + errMsg)
	}
}

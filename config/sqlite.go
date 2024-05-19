package config

import (
	"os"

	"github.com/GuilhermeHRC/go-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {

	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if database already exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("sqlite: database does not exist, creating it")

		// Create database file
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite: failed to connect database: %v", err)

		return nil, err
	}

	// Migrate schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite: failed to migrate schema: %v", err)

		return nil, err
	}

	// Return DB
	return db, nil

}

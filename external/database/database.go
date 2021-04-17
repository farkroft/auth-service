package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.com/farkroft/auth-service/external/config"
	"gitlab.com/farkroft/auth-service/external/constants"
	"gitlab.com/farkroft/auth-service/internal/model"
)

// Database struct
type Database struct {
	*gorm.DB
}

// NewDatabase init new database
func NewDatabase(cfg *config.Config) (*Database, error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.GetString(constants.EnvDBHost),
		cfg.GetString(constants.EnvDBPort),
		cfg.GetString(constants.EnvDBUser),
		cfg.GetString(constants.EnvDBPass),
		cfg.GetString(constants.EnvDBName))
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		return &Database{db}, err
	}

	return &Database{db}, err
}

// Close to close connection
func (d *Database) Close() error {
	err := d.Close()
	return err
}

// Migrate table
func (d *Database) Migrate() error {
	db := d.Debug().AutoMigrate(&model.User{})
	if db.Error != nil {
		return db.Error
	}

	return nil
}

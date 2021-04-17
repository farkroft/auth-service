package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"gitlab.com/farkroft/auth-service/application/request"
	"gitlab.com/farkroft/auth-service/external/database"
	"gitlab.com/farkroft/auth-service/internal/model"
)

// UserRepository interface
type UserRepository interface {
	RegisterUser(req request.UserRequest) (model.User, error)
	GetUser(query request.UserRequest) (model.User, error)
}

// UserRepo struct
type UserRepo struct {
	DB *gorm.DB
}

// NewUserRepository func
func NewUserRepository(db *database.Database) *UserRepo {
	return &UserRepo{DB: db.DB}
}

// RegisterUser record
func (u *UserRepo) RegisterUser(req request.UserRequest) (model.User, error) {
	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}
	db := u.DB.Create(&user)
	if db.Error != nil {
		return user, db.Error
	}
	return user, nil
}

// GetUser get user
func (u *UserRepo) GetUser(req request.UserRequest) (model.User, error) {
	query := model.User{
		Username: req.Username,
	}

	user := model.User{}
	db := u.DB.Debug().Where(query).First(&user)
	if db.Error != nil && !db.RecordNotFound() {
		return user, db.Error
	}

	if db.RecordNotFound() {
		return user, fmt.Errorf("record not found")
	}

	return user, nil
}

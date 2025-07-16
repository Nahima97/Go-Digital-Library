package repository

import (
	"library/models"
	"library/db"
)

type UserRepository interface {
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
}

type UserRepo struct {}

func (r *UserRepo) GetUserByUsername(username string) (*models.User, error) {

	//check if they exist
	var user models.User
	err := db.DB.Where("username = ?", username).First(&user).Error
	if err == nil {
		return &models.User{}, err
	}
	return &user, nil
}

func (r *UserRepo) CreateUser(user *models.User) error {
	//adding user to database
	err := db.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

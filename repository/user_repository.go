package repository

import (
	"library/db"
	"library/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRepository interface
type UserRepository interface {
GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
    GetUserByID(id uuid.UUID) (*models.User, error)
}

// Struct that implements the interface
type UserRepo struct {
    Db *gorm.DB
}

// Implementation of the interface
func (r *UserRepo) GetUserByID(id uuid.UUID) (*models.User, error) {
    var user models.User
    // Preload borrowed books if you want to access them directly
    if err := r.Db.Preload("BorrowedBooks").First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepo) GetUserByUsername(username string) (*models.User, error) {

	//check if they exist
	var user models.User
	err := db.Db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}

func (r *UserRepo) CreateUser(user *models.User) error {
	//adding user to database
	err := db.Db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil

}

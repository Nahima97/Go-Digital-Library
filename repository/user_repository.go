package repository

import (
	"library/db"
	"library/models"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error) 
}

type UserRepo struct {
	Db *gorm.DB
}

func (r *UserRepo) CreateUser(user *models.User) error {
	err := db.Db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}

func (r *UserRepo) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := db.Db.Preload("BorrowedBooks", "returned_at IS NULL").Preload("BorrowedBooks.Book").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err 
	}
	return &user, nil 
}


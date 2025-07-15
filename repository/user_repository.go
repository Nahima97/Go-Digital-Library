package repository

import (
    "go-digital/models"
    "gorm.io/gorm"
)

// UserRepository interface
type UserRepository interface {
    GetUserByID(id uint) (*models.User, error)
}

// Struct that implements the interface
type UserRepo struct {
    Db *gorm.DB
}

// Implementation of the interface
func (r *UserRepo) GetUserByID(id uint) (*models.User, error) {
    var user models.User
    // Preload borrowed books if you want to access them directly
    if err := r.Db.Preload("BorrowedBooks").First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

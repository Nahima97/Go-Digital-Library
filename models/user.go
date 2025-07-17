package models

import (

	"github.com/google/uuid"
)

type User struct {

	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username      string    `json:"username"`
	Email         string    `gorm:"unique"`
	Password      string    `json:"password"`
	Age           int       `json:"age"`
	UserRole      string    `json:"user_role"`
	BorrowedBooks []Book    `gorm:"many2many:user_book;" json:"borrowed_books"`
}

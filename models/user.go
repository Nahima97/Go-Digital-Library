package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name          string    `json:"name"`
	Email         string    `gorm:"unique"`
	Password      string    `json:"password"`
	Role          string    `json:"role"`
	BorrowedBooks []Loan    `gorm:"foreignKey:UserID" json:"borrowed_books"`
}

package models

import (
	"library/models"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title     string        `json:"title"`
	Author    string        `json:"author"`
	Year      int           `json:"year"`
	Genre     string        `json:"genre"`
	AgeRating int           `json:"age_rating"`
	UserID    []models.User `gorm:"many2many:user_book;"`
}

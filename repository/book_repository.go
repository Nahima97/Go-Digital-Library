package repository

import (
	"library/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetBookByTitle(title string) (models.Book, error)
	CreateBook(book models.Book) error
}

type bookRepo struct {
	db *gorm.DB
}

// GetBookByTitle fetches a book by title
func (r *bookRepo) GetBookByTitle(title string) (models.Book, error) {
	var book models.Book
	err := r.db.Where("title = ?", title).First(&book).Error
	return book, err
}

// CreateBook adds a new book
func (r *bookRepo) CreateBook(book models.Book) error {
	return r.db.Create(&book).Error
}

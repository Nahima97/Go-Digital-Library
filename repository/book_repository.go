package repository

import (
    "go-digital/models"
    "gorm.io/gorm"
)

// Interface for book-related DB operations
type BookRepository interface {
    GetBookByID(id uint) (*models.Book, error)
}

// Struct that implements the interface
type BookRepo struct {
    Db *gorm.DB
}

// GetBookByID fetches a book by ID
func (r *BookRepo) GetBookByID(id uint) (*models.Book, error) {
    var book models.Book
    // Preload users if you need to show who borrowed it (optional)
    if err := r.Db.First(&book, id).Error; err != nil {
        return nil, err
    }
    return &book, nil
}

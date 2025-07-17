package repository

import (
    "go-digital/models"
    "gorm.io/gorm"
  	"library/db"
)

// Interface for book-related DB operations
type BookRepository interface {
    GetBookByID(id uint) (*models.Book, error)
    DeleteBookByID(id uint) error
  	CreateBook(book models.Book) error

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


func (r *BookRepo) DeleteBookByID(id uint) error {
    return r.Db.Delete(&models.Book{}, id).Error
}

func (r BookRepo) SearchBooks(title, author, genre string) ([]models.Book, error) {
    var books []models.Book
    query := db.DB.Model(&models.Book{})

    if title != "" {
        query = query.Where("title ILIKE ?", "%"+title+"%")
    }
    if author != "" {
        query = query.Where("author ILIKE ?", "%"+author+"%")
    }
    if genre != "" {
        query = query.Where("genre ILIKE ?", "%"+genre+"%")
    }

    err := query.Find(&books).Error
    if err != nil {
        return nil, err
    }

    return books, nil
}





// func (r *BookRepo) SearchTitle(books string) (*models.Book, error){

// 	var book models.Book
// 	err := db.DB.Where("title = ?", books).Find(&book).Error
// 	if err == nil {
// 		return &models.Book{}, err
// 	}
// 	return &book, nil
// }

// func (r *BookRepo) SearchAuthor(books string) (*models.Book, error){

// 	var book models.Book
// 	err := db.DB.Where("author = ?", books).Find(&book).Error
// 	if err == nil {
// 		return &models.Book{}, err
// 	}
// 	return &book, nil
// }

// func (r *BookRepo) SearchGenre(books string) (*models.Book, error){

// 	var book models.Book
// 	err := db.DB.Where("genre = ?", books).Find(&book).Error
// 	if err == nil {
// 		return &models.Book{}, err
// 	}
// 	return &book, nil
// }


// GetBookByTitle fetches a book by title
func (r *BookRepo) GetBookByTitle(title string) (models.Book, error) {
	var book models.Book
	err := r.db.Where("title = ?", title).First(&book).Error
	return book, err
}

// CreateBook adds a new book
func (r *BookRepo) CreateBook(book models.Book) error {
	return r.db.Create(&book).Error

}

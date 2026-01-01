package repository

import (
	"library/db"
	"library/models"
	"time"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book models.Book) error
	GetBookByTitle(title string) (*models.Book, error)
	GetBookByID(id uuid.UUID) (*models.Book, error)
	GetBooks(title, author, genre string) ([]models.Book, error)
	BorrowBook(userID, bookID uuid.UUID) error
	UpdateBook(book *models.Book) (*models.Book, error)
	ReturnBook(userID, bookID uuid.UUID) error
	ArchiveBookByID(id uuid.UUID) error
}

type BookRepo struct {
	Db *gorm.DB
}

func (r *BookRepo) CreateBook(book models.Book) error {
	err := db.Db.Create(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *BookRepo) GetBookByTitle(title string) (*models.Book, error) {
	var book models.Book
	err := db.Db.Where("title = ?", title).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepo) GetBookByID(id uuid.UUID) (*models.Book, error) {
	var book models.Book
	err := db.Db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepo) GetBooks(title, author, genre string) ([]models.Book, error) {
	var books []models.Book
	query := db.Db.Model(&models.Book{}).Where("is_active = ?", true)

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

func (r *BookRepo) BorrowBook(userID, bookID uuid.UUID) error {
	dueDate := time.Now().Add(14 * 24 * time.Hour)

	loan := models.Loan{
		UserID:     userID,
		BookID:     bookID,
		BorrowedAt: time.Now(),
		DueDate:    &dueDate,
	}

    err := db.Db.Create(&loan).Error 
    if err != nil {
        return err 
    }
    return nil 
}

func (r *BookRepo) ReturnBook(userID, bookID uuid.UUID) error {
    var loan models.Loan
    err := db.Db.Where("user_id = ? AND book_id = ? AND returned_at IS NULL", userID, bookID).First(&loan).Error
    if err != nil {
        return err 
    }
    returnedAt := time.Now()
    loan.ReturnedAt = &returnedAt

    err = db.Db.Save(&loan).Error
    if err != nil {
        return err 
    }
    return nil 
}

func (r *BookRepo) UpdateBook(book *models.Book) (*models.Book, error) {
	err := db.Db.Model(&models.Book{}).Where("id = ?", book.ID).Updates(book).Error
	if err != nil {
		return nil, err
	}

	var updatedBook models.Book
	err = db.Db.Where("id = ?", book.ID).First(&updatedBook).Error
	if err != nil {
		return nil, err
	}
	return &updatedBook, nil
}

func (r *BookRepo) ArchiveBookByID(id uuid.UUID) error {
	err := db.Db.Model(&models.Book{}).Where("id = ?", id).UpdateColumn("is_active", false).Error
	if err != nil {
		return err
	}
	return nil
}

package services

import (
	"errors"
	"library/models"
	"library/repository"
	"net/http"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type BookService struct {
	Repo repository.BookRepository
}

func (s *BookService) AddBook(book models.Book) (models.Book, int, error) {
	// Basic validation
	if book.Title == "" || book.Author == "" {
		return models.Book{}, http.StatusBadRequest, errors.New("title and author are required")
	}

	existingBook, err := s.Repo.GetBookByTitle(book.Title)
	if err == nil && existingBook.ID != uuid.Nil {
		return models.Book{}, http.StatusBadRequest, errors.New("book already exists")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Book{}, http.StatusInternalServerError, errors.New("could not check for existing book")
	}

	// Set UUID as ID
	book.ID = uuid.New()

	// Create book
	if err := s.Repo.CreateBook(book); err != nil {
		return models.Book{}, http.StatusInternalServerError, errors.New("could not add book")
	}

	return book, http.StatusCreated, nil
}

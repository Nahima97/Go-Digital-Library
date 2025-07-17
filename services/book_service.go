package services

import (
	"errors"

	"library/models"
	"library/repository"
	"net/http"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookService struct {
	UserRepo repository.UserRepository
	BookRepo repository.BookRepository
	Db       *gorm.DB
}

func (s *BookService) ReturnBook(userID, bookID uuid.UUID) error {
	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	book, err := s.BookRepo.GetBookByID(bookID)
	if err != nil {
		return errors.New("book not found")
	}

	for _, b := range user.BorrowedBooks {
		if b.ID == book.ID {
			return errors.New("you have already borrowed this book")
		}
	}
	
	err = s.Db.Model(user).Association("BorrowedBooks").Delete(book) 
	if err != nil {
		return errors.New("failed to borrow book")
	}

	return nil
}

func (s *BookService) AddBook(book models.Book) (models.Book, int, error) {

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

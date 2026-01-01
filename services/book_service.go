package services

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"library/models"
	"library/repository"
)

type BookService struct {
	UserRepo repository.UserRepository
	BookRepo repository.BookRepository
}

func (s *BookService) AddBook(book models.Book) error {
	if book.Title == "" {
		return errors.New("book title is required")
	}

	if book.Author == "" {
		return errors.New("author is required")
	}

	if book.Description == "" {
		return errors.New("book description is required")
	}

	if book.Year == 0 {
		return errors.New("book year is required")
	}

	if book.Genre == "" {
		return errors.New("book genre is required")
	}

	existingBook, err := s.BookRepo.GetBookByTitle(book.Title)
	if err == nil && existingBook != nil {
		return fmt.Errorf("book with name %q already exists", book.Title)
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to check book existence: %w", err)
	}

	err = s.BookRepo.CreateBook(book)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookService) GetBooks(title, author, genre string) ([]models.Book, error) {
	books, err := s.BookRepo.GetBooks(title, author, genre)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *BookService) GetBooksByUserID(userID uuid.UUID) (*models.User, error) {
	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *BookService) BorrowBook(userID, bookID uuid.UUID) error {
	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	book, err := s.BookRepo.GetBookByID(bookID)
	if err != nil {
		return errors.New("book not found")
	}

	for _, b := range user.BorrowedBooks {
		if b.BookID == book.ID {
			return errors.New("you are already borrowing this book")
		}
	}

	err = s.BookRepo.BorrowBook(userID, bookID)
	if err != nil {
		return err
	}
	return nil
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
		if b.BookID == book.ID {
			return s.BookRepo.ReturnBook(userID, bookID)
		}
	}
	return errors.New("cannot return book as you are not borrowing it")
}

func (s *BookService) UpdateBook(bookID uuid.UUID, req models.Book) (*models.Book, error) {
	book, err := s.BookRepo.GetBookByID(bookID)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		book.Title = req.Title
	}

	if req.Author != "" {
		book.Author = req.Author
	}

	if req.Description != "" {
		book.Description = req.Description
	}

	if req.Year != 0 {
		book.Year = req.Year
	}

	if req.Genre != "" {
		book.Genre = req.Genre
	}

	updatedBook, err := s.BookRepo.UpdateBook(book)
	if err != nil {
		return nil, err
	}
	return updatedBook, nil
}

func (s *BookService) ArchiveBook(bookID uuid.UUID) error {
	err := s.BookRepo.ArchiveBookByID(bookID)
	if err != nil {
		return err
	}
	return nil
}

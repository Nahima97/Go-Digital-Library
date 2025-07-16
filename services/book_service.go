package services

import (
	"errors"
	"library/repository"

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

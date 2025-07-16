package services

import (
	"errors"
	"go-digital/models"
	"go-digital/repository"
	"gorm.io/gorm"
)

type BookService struct {
	UserRepo repository.UserRepository
	BookRepo repository.BookRepository
	Db       *gorm.DB // Required to use GORM's association features
}

func (s *BookService) BorrowBook(userID, bookID uint) error {
	// 1. Get the user
	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	// 2. Get the book
	book, err := s.BookRepo.GetBookByID(bookID)
	if err != nil {
		return errors.New("book not found")
	}

	// 3. Check if user is old enough
	if user.Age < book.AgeRating {
		return errors.New("user must be old enough to read books")
	}

	// 4. Check if the user already borrowed the book
	for _, b := range user.BorrowedBooks {
		if b.ID == book.ID {
			return errors.New("you have already borrowed this book")
		}
	}

	// 5. Append the book to user's BorrowedBooks (many-to-many link)
	if err := s.Db.Model(user).Association("BorrowedBooks").Append(book); err != nil {
		return errors.New("failed to borrow book")
	}

	return nil
}

func (s *BookService) DeleteBook(userID, bookID uint) error {
    user, err := s.UserRepo.GetUserByID(userID)
    if err != nil {
        return errors.New("user not found")
    }

    if user.UserRole != "admin" {
        return errors.New("admin access only")
    }

    return s.BookRepo.DeleteBookByID(bookID)
}

}

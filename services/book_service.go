package services

import (
	"library/models"
	"library/repository"
)


type BookService struct {
	Repo repository.BookRepository
}

func (s *BookService) SearchBooks(title, author, genre string) ([]models.Book, error) {

	books, err := s.Repo.SearchBooks(title, author, genre)
    if err != nil {
        return nil, err
    }
    return books, nil
	
}


// func (s *BookService) SearchTitle(req *models.Book) error {
// 	_, err := s.Repo.SearchTitle(req.Title)
// 	if err != nil {
// 		return err
// 	}
// 	return nil 
// }

// func (s *BookService) SearchAuthor(req *models.Book) error {
// 	_, err := s.Repo.SearchAuthor(req.Author)
// 	if err != nil {
// 		return err
// 	}
// 	return nil 
// }

// func (s *BookService) SearchGenre(req *models.Book) error {
// 	_, err := s.Repo.SearchGenre(req.Genre)
// 	if err != nil {
// 		return err
// 	}
// 	return nil 
// }
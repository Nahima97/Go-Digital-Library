package repository

import	(		 
	"library/models"	
	"library/db"
)

type BookRepository interface {

	SearchBooks(title, author, genre string) ([]models.Book, error)
	
}


type BookRepo struct {}

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
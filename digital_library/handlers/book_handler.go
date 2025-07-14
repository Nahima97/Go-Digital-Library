package handlers

import "library/services"

type BookHandler struct {
	Service *services.BookService
}


func () SearchBook() {
//all - guests, users


}

func () BorrowBook() {
//users
//check if user is old enough
//delete from the users slice of books 

}

func () ReturnBook() {
//users
//append the users slice of books 

}

func () AddBook() {
//only admin 


}

func () DeleteBook() {
//only admin


}
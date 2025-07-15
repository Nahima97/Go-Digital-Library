package handlers

import "library/services"

type BookHandler struct {
	Service *services.BookService
}


func () SearchBook() {
//all - guests, users

//joshua and pana
}

func () BorrowBook() {
//users
//check if user is old enough
//delete from the users slice of books 

//nahima, alina, safa
}

func () ReturnBook() {
//users
//append the users slice of books 


//nahima, alina, safa
}

func () AddBook() {
//only admin 


//nahima, alina, safa
}

func () DeleteBook() {
//only admin


//nahima, alina, safa
}
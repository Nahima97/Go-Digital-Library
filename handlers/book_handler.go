package handlers

import (
	"encoding/json"
	"library/services"
	"net/http"
)

type BookHandler struct {
	Service *services.BookService
}


func (h *BookHandler) SearchTitle(w http.ResponseWriter, r *http.Request) {
	// all - guests, users

	title := r.URL.Query().Get("title")
	author := r.URL.Query().Get("author")
	genre := r.URL.Query().Get("genre")

	books, err := h.Service.SearchBooks(title, author, genre)
	if err != nil {
		http.Error(w, "Error searching books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
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
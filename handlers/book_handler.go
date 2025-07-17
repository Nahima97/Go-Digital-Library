package handlers

import (
	"encoding/json"
	"library/db"
	"library/models"
	"library/services"
	"net/http"
)

type BookHandler struct {
	Service *services.BookService
}

func () SearchBook() {
//all - guests, users

}

func BorrowBook() {
	//users
	//check if user is old enough
	//delete from the users slice of books

}

func ReturnBook() {
	//users
	//append the users slice of books

}

func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	
	userClaims := r.Context().Value(middleware.UserKey).(*services.Claims)

	if !userClaims.IsAdmin {
		http.Error(w, "Access denied: admin only", http.StatusForbidden)
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	addedBook, status, err := h.Service.AddBook(book)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(addedBook)
}

func DeleteBook() {
	//only admin

}

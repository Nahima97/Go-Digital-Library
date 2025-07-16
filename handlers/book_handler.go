package handlers

import (
	"library/services"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

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

func (h *BookHandler) ReturnBook(w http.ResponseWriter, r *http.Request) {
 userID, err := utils.ExtractUserID(r)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
	
bookID := strings.TrimPrefix(r.URL.Path, "/return/:")

bookUUID, _ := uuid.Parse(bookID)
	
      err = h.Service.ReturnBook(uuid.UUID(userID), uuid.UUID(bookUUID))
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

w.WriteHeader(http.StatusOK)
}

func () AddBook() {
//only admin 


//nahima, alina, safa
}

func () DeleteBook() {
//only admin


//nahima, alina, safa
}
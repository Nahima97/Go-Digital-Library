package handlers

import (
	"library/services"
	"net/http"
	"strconv"
)

type BookHandler struct {
	Service *services.BookService
}


func () SearchBook() {
//all - guests, users


}

func (h *BookHandler) BorrowBook(w http.ResponseWriter, r *http.Request) {
    userID, err := utils.ExtractUserID(r)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    bookIDStr := r.URL.Query().Get("id")
    if bookIDStr == "" {
        http.Error(w, "Missing book ID", http.StatusBadRequest)
        return
    }

    bookID, err := strconv.Atoi(bookIDStr)
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }

    err = h.Service.BorrowBook(uint(userID), uint(bookID))
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Book borrowed successfully",
    })
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
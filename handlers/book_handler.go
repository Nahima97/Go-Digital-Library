package handlers

import (
	"strconv"

	"strings"

	"github.com/google/uuid"
	"encoding/json"
	"library/db"
	"library/models"
	"library/services"
	"net/http"

)

type BookHandler struct {
	Service *services.BookService
}



func (h *BookHandler) SearchTitle(w http.ResponseWriter, r *http.Request) {
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


func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
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

    err = h.Service.DeleteBook(uint(userID), uint(bookID))
    if err != nil {
        http.Error(w, err.Error(), http.StatusForbidden)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Book deleted successfully",
    })
}


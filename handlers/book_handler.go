package handlers

import (
	"strings"
	"encoding/json"
	"library/config"
	"library/models"
	"library/services"
	"library/utils"
	"net/http"

	"github.com/google/uuid"
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
	userID, err := utils.ExtractUserID(w, r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userUUID, _ := uuid.Parse(userID)

	bookID := strings.TrimPrefix(r.URL.Path, "/return/:")

	bookUUID, _ := uuid.Parse(bookID)

	err = h.Service.BorrowBook(uuid.UUID(userUUID), uuid.UUID(bookUUID))
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
	userID, err := utils.ExtractUserID(w, r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userUUID, _ := uuid.Parse(userID)

	bookID := strings.TrimPrefix(r.URL.Path, "/return/:")

	bookUUID, _ := uuid.Parse(bookID)

	err = h.Service.ReturnBook(uuid.UUID(userUUID), uuid.UUID(bookUUID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

    userRole, err := utils.GetUserRole(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	permission := config.RolePermission(userRole, "add:book")
	if !permission {
		http.Error(w, "access denied", http.StatusUnauthorized)
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
	userRole, err := utils.GetUserRole(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	permission := config.RolePermission(userRole, "add:book")
	if !permission {
		http.Error(w, "access denied", http.StatusUnauthorized)
		return
	}

    userID, err := utils.ExtractUserID(w, r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userUUID, _ := uuid.Parse(userID)

	bookID := strings.TrimPrefix(r.URL.Path, "/return/:")

	bookUUID, _ := uuid.Parse(bookID)

	err = h.Service.DeleteBook(uuid.UUID(userUUID), uuid.UUID(bookUUID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Book deleted successfully",
	})
}

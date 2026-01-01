package handlers

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"library/middleware"
	"library/models"
	"library/services"
	"net/http"
)

type BookHandler struct {
	Service *services.BookService
}

func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	role := middleware.GetUserRole(r.Context())
	if role != "admin" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return 
	}

	err = h.Service.AddBook(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	author := r.URL.Query().Get("author")
	genre := r.URL.Query().Get("genre")

	books, err := h.Service.GetBooks(title, author, genre)
	if err != nil {
		http.Error(w, "unable to get books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) GetBooksByUserID(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	if userID == uuid.Nil {
		http.Error(w, "invalid user ID", http.StatusUnauthorized)
		return
	}

	user, err := h.Service.GetBooksByUserID(userID)
	if err != nil {
		http.Error(w, "unable to get books for user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user.BorrowedBooks)
}

func (h *BookHandler) BorrowBook(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	if userID == uuid.Nil {
		http.Error(w, "invalid user ID", http.StatusUnauthorized)
		return
	}

	bookID := mux.Vars(r)["id"]
	bookUUID, err := uuid.FromString(bookID)
	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}

	err = h.Service.BorrowBook(userID, bookUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("book borrowed successfully"))
}

func (h *BookHandler) ReturnBook(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	if userID == uuid.Nil {
		http.Error(w, "invalid user ID", http.StatusUnauthorized)
		return
	}

	bookID := mux.Vars(r)["id"]
	bookUUID, err := uuid.FromString(bookID)
	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}

	err = h.Service.ReturnBook(userID, bookUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("book returned successfully"))
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var req models.Book
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	role := middleware.GetUserRole(r.Context())
	if role != "admin" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	bookID := mux.Vars(r)["id"]
	bookUUID, err := uuid.FromString(bookID)
	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}

	updatedBook, err := h.Service.UpdateBook(bookUUID, req)
	if err != nil {
		http.Error(w, "unable to update book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBook)
}

func (h *BookHandler) ArchiveBook(w http.ResponseWriter, r *http.Request) {
	role := middleware.GetUserRole(r.Context())
	if role != "admin" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	bookID := mux.Vars(r)["id"]
	bookUUID, err := uuid.FromString(bookID)
	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}

	err = h.Service.ArchiveBook(bookUUID)
	if err != nil {
		http.Error(w, "unable to archive book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("book archived successfully"))
}

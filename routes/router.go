package routes

import (
	"library/handlers"
	"library/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handlers.UserHandler, bookHandler *handlers.BookHandler) *mux.Router {

	r := mux.NewRouter()
	// Health check or root
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Library API"))
	})

	// user routes

	r.HandleFunc("/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/register", userHandler.Register).Methods("POST")

	// Book Routes
	r.HandleFunc("/books", bookHandler.SearchTitle).Methods("GET")

	// protected routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	protected.HandleFunc("/add", bookHandler.AddBook).Methods("POST")
	protected.HandleFunc("/books/borrow/{id}", bookHandler.BorrowBook).Methods("GET")
	protected.HandleFunc("/books/return/{id}", bookHandler.ReturnBook).Methods("PUT")
	protected.HandleFunc("/books/delete/{id}", bookHandler.DeleteBook).Methods("DELETE")

	return r
}

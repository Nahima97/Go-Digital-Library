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

	r.HandleFunc("/login", userHandler.Login)
	r.HandleFunc("/register", userHandler.Register)

	// Book Routes
	r.HandleFunc("/books", bookHandler.SearchTitle)

	// protected routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	protected.HandleFunc("/add", bookHandler.AddBook)
	protected.HandleFunc("/{id}", bookHandler.BorrowBook)
	protected.HandleFunc("/{id}", bookHandler.ReturnBook)
	protected.HandleFunc("/{id}", bookHandler.DeleteBook)

	return r
}

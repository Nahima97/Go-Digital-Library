package routes

import (
	"library/handlers"
	"library/middleware"
	"net/http"
	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handlers.UserHandler, bookHandler *handlers.BookHandler) *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Library API"))
	})

	r.HandleFunc("/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/books", bookHandler.GetBooks).Methods("GET")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	protected.HandleFunc("/add", bookHandler.AddBook).Methods("POST")
	protected.HandleFunc("/books/user", bookHandler.GetBooksByUserID).Methods("GET") 
	protected.HandleFunc("/books/{id}/borrow", bookHandler.BorrowBook).Methods("GET")
	protected.HandleFunc("/books/{id}/return", bookHandler.ReturnBook).Methods("PUT")
	protected.HandleFunc("/books/{id}/update", bookHandler.UpdateBook).Methods("PATCH")
	protected.HandleFunc("/books/{id}/archive", bookHandler.ArchiveBook).Methods("PATCH")

	return r
}

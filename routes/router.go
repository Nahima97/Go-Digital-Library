package routes

import (
	"go-digital/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handlers.UserHandler, bookHandler *handlers.BookHandler) *mux.Router {
	r := mux.NewRouter()

	// Route for borrowing a book
	r.HandleFunc("/borrow", bookHandler.BorrowBook).Methods("POST")

	return r
}

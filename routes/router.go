package routes

import (
	"go-digital/handlers"
	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handlers.UserHandler, bookHandler *handlers.BookHandler) *mux.Router {
	r := mux.NewRouter()

//mosun 
  r.HandleFunc("/borrow", bookHandler.BorrowBook).Methods("POST")



}


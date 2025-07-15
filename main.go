package main

import (
	"library/db"
	"library/handlers"
	"library/repository"
	"library/routes"
	"library/services"
)

func main() {

	db.InitDb()

	userRepo := &repository.UserRepo{}
	bookRepo := &repository.BookRepo{}

	userService := &services.UserService{Repo: userRepo}
	bookService := &services.BookService{Repo: bookRepo}

	userHandler := &handlers.UserHandler{Service: userService}
	bookHandler := &handlers.BookHandler{Service: bookService}

	r := routes.SetupRouter(userHandler, bookHandler)

	//start server

}

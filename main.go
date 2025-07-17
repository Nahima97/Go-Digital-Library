package main

import (
	"fmt"
	"library/db"
	"library/handlers"
	"library/repository"
	"library/routes"
	"library/services"
	"log"
	"net/http"
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

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("failed to start server", err)
	}
	fmt.Println("server started!")
}

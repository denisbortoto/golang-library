package main

import (
	"fmt"
	"library/internal/controller"
	"library/internal/repository"

	"library/internal/service"
	"library/internal/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db := utils.Connect()
	defer db.Close()
	r := mux.NewRouter()

	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)
	bookController.Routes(r)

	port := "8000"
	fmt.Printf("Api running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

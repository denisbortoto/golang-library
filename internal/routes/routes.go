package routes

import (
	"library/internal/controller"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Home).Methods("GET")
	r.HandleFunc("/books", controller.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controller.GetBook).Methods("GET")
	r.HandleFunc("/books", controller.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", controller.UpdateBook).Methods("PATCH")
	return r
}

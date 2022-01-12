package controller

import (
	"encoding/json"
	"library/internal/entity"
	"library/internal/errors"
	"library/internal/service"
	"library/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var (
	bookService service.BookService = service.NewBookService()
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	/* GetBooks function is used to communicate with the service layer to validate the business logic
	and return all books listed in the application */
	books, err := bookService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting books"})
		return
	}
	utils.ToJson(w, books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	/* GetBook function is used to communicate with the service layer to validate the business logic
	and return the book specified with the ID */
	params := mux.Vars(r)
	book_id, err := uuid.Parse(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Invalid ID"})
		return
	}

	result, err := bookService.GetById(book_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Book not found"})
		return
	}
	utils.ToJson(w, result)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	/* CreateBook function is used to communicate with the service layer to validate the business logic
	in order to insert a new book in the database */

	// body, _ := ioutil.ReadAll(r.Body)
	//var book entity.Book
	//err := json.Unmarshal(body, &book)

	var book entity.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	err1 := bookService.Validate(&book)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := bookService.Create(&book)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error creating the book"})
		return
	}
	utils.ToJson(w, result)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	/* UpdateBook function is used to communicate with the service layer to validate the business logic
	in order to update information on a existing book in the database */
	params := mux.Vars(r)
	book_id, err := uuid.Parse(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Invalid ID"})
		return
	}

	var book entity.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	book.Id = book_id
	result, err := bookService.Update(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error updating the book"})
	}

	utils.ToJson(w, result)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	/* DeleteBook function is used to communicate with the service layer to validate the business logic
	in order to delete a existing book from the database */
	params := mux.Vars(r)
	book_id, err := uuid.Parse(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Invalid ID"})
		return
	}

	result, err := bookService.Delete(book_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error deleting the book"})
	}

	utils.ToJson(w, result)
}

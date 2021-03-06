package repository

import (
	"library/internal/entity"
	"library/internal/utils"

	"github.com/google/uuid"
)

type BookRepository interface {
	Save(book *entity.Book) (*entity.Book, error)
	Update(book entity.Book) (int64, error)
	GetAll() ([]entity.Book, error)
	GetById(book_id uuid.UUID) (entity.Book, error)
	Delete(book_id uuid.UUID) (int64, error)
}

type bookRepo struct{}

//NewBookRepository
func NewBookRepository() BookRepository {
	return &bookRepo{}
}

func (*bookRepo) Save(book *entity.Book) (*entity.Book, error) {
	/* Save function is used to connect to the DB in order to exec the query to save a new book to the DB */
	db := utils.Connect()
	defer db.Close()

	stmt := (`INSERT INTO books (book_id, title, description, isbn, price) VALUES ($1, $2, $3, $4, $5)`)
	book.Id = uuid.New()
	_, err := db.Query(stmt, book.Id, book.Title, book.Description, book.Isbn, book.Price)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (*bookRepo) Update(book entity.Book) (int64, error) {
	/* Update function is used to connect to the DB in order to exec the query to update a existing book in the DB */

	db := utils.Connect()
	defer db.Close()

	stmt := (`UPDATE books SET title = $1, description = $2, isbn = $3, price = $4 WHERE book_id = $5`)
	result, err := db.Exec(stmt, book.Title, book.Description, book.Isbn, book.Price, book.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (*bookRepo) GetById(book_id uuid.UUID) (entity.Book, error) {
	/* GetById function is used to connect to the DB in order to exec the query that finds a existing book in the DB by its ID */

	db := utils.Connect()
	defer db.Close()
	stmt := (`SELECT * FROM books WHERE book_id = $1`)
	result, err := db.Query(stmt, book_id)
	if err != nil {
		return entity.Book{}, err
	}
	defer result.Close()
	var book entity.Book
	for result.Next() {
		err = result.Scan(&book.Id, &book.Title, &book.Description, &book.Isbn, &book.Price)
		if err != nil {
			return entity.Book{}, err
		}
	}
	return book, err
}

func (*bookRepo) GetAll() ([]entity.Book, error) {
	/* GetAll function is used to connect to the DB in order to exec the query that finds all existing books in the DB */

	db := utils.Connect()
	defer db.Close()
	stmt := (`SELECT * FROM books`)
	result, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var books []entity.Book
	for result.Next() {
		var book entity.Book
		err := result.Scan(&book.Id, &book.Title, &book.Description, &book.Isbn, &book.Price)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (*bookRepo) Delete(book_id uuid.UUID) (int64, error) {
	/* Delete function is used to connect to the DB in order to exec the query to delete a existing book in the DB */

	db := utils.Connect()
	defer db.Close()

	stmt := (`DELETE from books WHERE book_id=$1`)
	result, err := db.Exec(stmt, book_id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

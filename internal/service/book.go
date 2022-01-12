package service

import (
	"errors"
	"library/internal/entity"
	"library/internal/repository"

	"github.com/google/uuid"
)

type service struct{}

type BookService interface {
	Validate(book *entity.Book) error
	Create(book *entity.Book) (*entity.Book, error)
	Update(book entity.Book) (int64, error)
	GetById(book_id uuid.UUID) (entity.Book, error)
	GetAll() ([]entity.Book, error)
	Delete(book_id uuid.UUID) (int64, error)
}

var (
	repo repository.BookRepository = repository.NewBookRepository()
)

func NewBookService() BookService {
	return &service{}
}

func (*service) Validate(book *entity.Book) error {
	/* Validate function contains the business logic utilized to create a new book */
	if book.Title == "" {
		err := errors.New("Book Title cannot be empty")
		return err
	}
	if book.Description == "" {
		err := errors.New("Book Description cannot be empty")
		return err
	}
	if book.Isbn == "" {
		err := errors.New("Book Isbn cannot be empty")
		return err
	}
	if book.Price == 0 {
		err := errors.New("Book price cannot be empty")
		return err
	}
	return nil
}

func (*service) Create(book *entity.Book) (*entity.Book, error) {
	return repo.Save(book)
}

func (*service) Update(book entity.Book) (int64, error) {
	return repo.Update(book)
}

func (*service) GetById(book_id uuid.UUID) (entity.Book, error) {
	return repo.GetById(book_id)
}

func (*service) GetAll() ([]entity.Book, error) {
	return repo.GetAll()
}

func (*service) Delete(book_id uuid.UUID) (int64, error) {
	return repo.Delete(book_id)
}

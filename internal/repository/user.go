package repository

import (
	"library/internal/entity"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	Update(user entity.User) (int64, error)
	GetAll() ([]entity.User, error)
	GetById(user_id uuid.UUID) (entity.User, error)
}

//type userRepo struct{}

//NewUserRepository
//func NewUserRepository() UserRepository {
//	return &userRepo{}
//}

/* func (*userRepo) Save(user *entity.User) (*entity.User, error) {
	db := utils.Connect()
	defer db.Close()

	stmt := (`INSERT INTO book(book_id, title, description, isbn, price) VALUES ($1, $2, $3, $4, $5)`)
	book.Id = uuid.New()
	_, err := db.Query(stmt, book.Id, book.Title, book.Description, book.Isbn, book.Price)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (*userRepo) Update(book entity.Book) (int64, error) {
	db := utils.Connect()
	defer db.Close()

	stmt := (`UPDATE book SET title = $1, description = $2, isbn = $3, price = $4 WHERE book_id = $5`)
	result, err := db.Exec(stmt, book.Title, book.Description, book.Isbn, book.Price, book.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (*userRepo) GetById(book_id uuid.UUID) (entity.Book, error) {
	db := utils.Connect()
	defer db.Close()
	stmt := (`SELECT * FROM book WHERE book_id = $1`)
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

func (*userRepo) GetAll() ([]entity.Book, error) {
	db := utils.Connect()
	defer db.Close()
	stmt := (`SELECT * FROM book`)
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
*/

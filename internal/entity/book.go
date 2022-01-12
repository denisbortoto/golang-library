package entity

import "github.com/google/uuid"

type Book struct {
	Id          uuid.UUID `json:"book_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Isbn        string    `json:"isbn"`
	Price       float64   `json:"price"`
}

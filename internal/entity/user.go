package entity

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID `json:"user_id"`
	Name      string    `json:"first_name"`
	Last_name string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Cpf       string    `json:"cpf"`
}

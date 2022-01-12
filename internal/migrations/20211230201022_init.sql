-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
    user_id uuid DEFAULT uuid_generate_v4(),
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    email   VARCHAR NOT NULL UNIQUE,
    password VARCHAR (15) NOT NULL UNIQUE,
    cpf  VARCHAR(11) NOT NULL UNIQUE,
    PRIMARY KEY (user_id)
);
CREATE TABLE books (
    book_id uuid DEFAULT uuid_generate_v4(),
    title VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    isbn   VARCHAR NOT NULL,
    Price  float NOT NULL,
    PRIMARY KEY (book_id)
);

-- +goose Down
DROP TABLE users;
DROP TABLE books;


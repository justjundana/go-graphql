package book

import (
	_models "github.com/justjundana/go-graphql/models"
)

type BookInterface interface {
	GetBooks() ([]_models.Book, error)
	GetBook(id int) (_models.Book, error)
	CreateBook(book _models.Book) (_models.Book, error)
	UpdateBook(book _models.Book) (_models.Book, error)
	DeleteBook(book _models.Book) (_models.Book, error)
}

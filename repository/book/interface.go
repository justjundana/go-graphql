package book

import (
	_models "github.com/justjundana/go-graphql/models"
)

type BookInterface interface {
	GetBooks() ([]_models.Book, error)
	GetBook(id int) (_models.Book, error)
}

package book

import (
	"database/sql"
	"log"

	_models "github.com/justjundana/go-graphql/models"
)

type BookRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (br *BookRepository) GetBooks() ([]_models.Book, error) {
	var books []_models.Book
	rows, err := br.db.Query(`SELECT id, title, description, author, publisher FROM books ORDER BY id ASC`)
	if err != nil {
		log.Fatalf("Error")
	}

	defer rows.Close()

	for rows.Next() {
		var book _models.Book

		err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.Author, &book.Publisher)
		if err != nil {
			log.Fatalf("Error")
		}

		books = append(books, book)
	}

	return books, nil
}

func (br *BookRepository) GetBook(id int) (_models.Book, error) {
	var book _models.Book

	row := br.db.QueryRow(`SELECT id, title, description, author, publisher, status FROM books WHERE id = ?`, id)

	err := row.Scan(&book.ID, &book.Title, &book.Description, &book.Author, &book.Publisher, &book.Status)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (br *BookRepository) CreateBook(book _models.Book) (_models.Book, error) {
	query := `INSERT INTO books (title, description, author, publisher, status) VALUES (?, ?, ?, ?, ?)`

	statement, err := br.db.Prepare(query)
	if err != nil {
		return book, err
	}

	defer statement.Close()

	_, err = statement.Exec(book.Title, book.Description, book.Author, book.Publisher, book.Status)
	if err != nil {
		return book, err
	}

	return book, nil
}

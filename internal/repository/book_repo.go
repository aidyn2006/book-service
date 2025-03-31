package repository

import (
	"book-service/internal/entity"
	"database/sql"
	"fmt"
)

type BookRepository interface {
	GetByID(id int64) (*entity.Book, error)
	Create(book *entity.Book) error
}

type bookRepo struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepo{db}
}

func (r *bookRepo) GetByID(id int64) (*entity.Book, error) {
	var book entity.Book
	query := `SELECT id, title, author, published FROM books WHERE id = $1`
	err := r.db.QueryRow(query, id).
		Scan(&book.ID, &book.Title, &book.Author, &book.Published)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("book not found")
		}
		return nil, err
	}
	return &book, nil
}

func (r *bookRepo) Create(book *entity.Book) error {
	query := `INSERT INTO books (title, author, published) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, book.Title, book.Author, book.Published)
	return err
}

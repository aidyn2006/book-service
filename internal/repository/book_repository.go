package repository

import (
	"book-service/internal/entity"
)

type BookRepository interface {
	GetByID(id int64) (*entity.Book, error)
	Create(book *entity.Book) (*entity.Book, error)
	Update(book *entity.Book) (*entity.Book, error)
	Delete(id int64) error
	CheckAvailability(id int64) (bool, error)
	UpdateStock(id int64, quantity int) error
}

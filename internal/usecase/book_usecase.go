package usecase

import (
	"book-service/internal/entity"
	"book-service/internal/repository"
)

type BookUsecase interface {
	GetBook(id int64) (*entity.Book, error)
	AddBook(book *entity.Book) error
}

type bookUsecase struct {
	repo repository.BookRepository
}

func NewBookUsecase(repo repository.BookRepository) BookUsecase {
	return &bookUsecase{repo}
}

func (u *bookUsecase) GetBook(id int64) (*entity.Book, error) {
	return u.repo.GetByID(id)
}

func (u *bookUsecase) AddBook(book *entity.Book) error {
	return u.repo.Create(book)
}

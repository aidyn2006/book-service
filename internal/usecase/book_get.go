package usecase

import (
	"book-service/internal/entity"
	"book-service/internal/repository"
)

type BookGetter interface {
	GetBook(id int64) (*entity.Book, error)
}

type bookGetter struct {
	repo repository.BookRepository
}

func NewBookGetter(repo repository.BookRepository) BookGetter {
	return &bookGetter{repo}
}

func (u *bookGetter) GetBook(id int64) (*entity.Book, error) {
	return u.repo.GetByID(id)
}

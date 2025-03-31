package usecase

import "book-service/internal/repository"

type BookAvailabilityChecker interface {
	CheckAvailability(id int64) (bool, error)
}

type bookAvailabilityChecker struct {
	repo repository.BookRepository
}

func NewBookAvailabilityChecker(repo repository.BookRepository) BookAvailabilityChecker {
	return &bookAvailabilityChecker{repo}
}

func (u *bookAvailabilityChecker) CheckAvailability(id int64) (bool, error) {
	return u.repo.CheckAvailability(id)
}

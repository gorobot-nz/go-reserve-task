package usecase

import (
	"go-tech-task/internal/domain"
)

type BookUsecase interface {
	GetBooks() ([]domain.Book, error)
	GetBookById(id int64) (domain.Book, error)
	AddBooks(book *domain.Book) int64
	DeleteBook(id int64) (int64, error)
	UpdateBook(id int64) (int64, error)
}

type Usecase struct {
	BookUsecase
}


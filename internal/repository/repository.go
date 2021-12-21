package repository

import "go-tech-task/internal/domain"

type Book interface {
	GetBooks() ([]domain.Book, error)
	GetBookById(id int64) (domain.Book, error)
	AddBooks(book *domain.Book) int64
	DeleteBook(id int64) (int64, error)
	UpdateBook(id int64) (int64, error)
}

type Repository struct {
	Book
}

func NewRepository(books []domain.Book) *Repository {
	return &Repository{Book: NewLocalBook(books)}
}

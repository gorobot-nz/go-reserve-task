package domain

import (
	"github.com/lib/pq"

	"context"
)

type Book struct {
	ID      int64          `json:"id" db:"id"`
	Title   string         `json:"title" binding:"required" db:"title"`
	Authors pq.StringArray `json:"authors" binding:"required" db:"authors"`
	Year    string         `json:"year" binding:"required" db:"book_year"`
}

type BookRepository interface {
	GetBooks(ctx context.Context) ([]Book, error)
	GetBookById(ctx context.Context, id int64) (*Book, error)
	AddBooks(ctx context.Context, book Book) (int64, error)
	DeleteBook(ctx context.Context, id int64) (int64, error)
	UpdateBook(ctx context.Context, id int64, book Book) (int64, error)
}

type BookUseCase interface {
	GetBooks(ctx context.Context) ([]Book, error)
	GetBookById(ctx context.Context, id int64) (*Book, error)
	AddBooks(ctx context.Context, book Book) (int64, error)
	DeleteBook(ctx context.Context, id int64) (int64, error)
	UpdateBook(ctx context.Context, id int64, book Book) (int64, error)
}

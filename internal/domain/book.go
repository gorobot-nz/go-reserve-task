package domain

import (
	"context"
	"github.com/lib/pq"
)

type Book struct {
	ID      string         `json:"_id" db:"id"`
	Title   string         `json:"title" binding:"required" db:"title"`
	Authors pq.StringArray `json:"authors" db:"authors"`
	Year    string         `json:"year" binding:"required" db:"book_year"`
}

type BookRepository interface {
	GetBooks(ctx context.Context) ([]Book, error)
	GetBookById(ctx context.Context, id string) (*Book, error)
	AddBooks(ctx context.Context, book Book) (string, error)
	DeleteBook(ctx context.Context, id string) (string, error)
	UpdateBook(ctx context.Context, id string, book Book) (string, error)
}

type BookUseCase interface {
	GetBooks(ctx context.Context) ([]Book, error)
	GetBookById(ctx context.Context, id string) (*Book, error)
	AddBooks(ctx context.Context, book Book) (string, error)
	DeleteBook(ctx context.Context, id string) (string, error)
	UpdateBook(ctx context.Context, id string, book Book) (string, error)
}

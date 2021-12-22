package domain

import (
	"context"
	"time"
)

type Book struct {
	ID      int64    `json:"id" binding:"required"`
	Title   string   `json:"title" binding:"required"`
	Authors []string `json:"authors" binding:"required"`
	Year    time.Time   `json:"year" binding:"required" time_format:"2006-01-02T15:04:05Z07:00"`
}

type BookRepository interface {
	GetBooks(ctx context.Context) ([]Book, error)
	GetBookById(ctx context.Context, id int64) (Book, error)
	AddBooks(ctx context.Context, book Book) (int64, error)
	DeleteBook(ctx context.Context, id int64) (int64, error)
	UpdateBook(ctx context.Context, id int64, book Book) (int64, error)
}

type BookUseCase interface {
	GetBooks(ctx context.Context) ([]Book, error)
	GetBookById(ctx context.Context, id int64) (Book, error)
	AddBooks(ctx context.Context, book Book) (int64, error)
	DeleteBook(ctx context.Context, id int64) (int64, error)
	UpdateBook(ctx context.Context, id int64, book Book) (int64, error)
}

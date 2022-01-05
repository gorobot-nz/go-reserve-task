package domain

import (
	"context"
)

type Book struct {
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Year    string   `json:"year"`
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

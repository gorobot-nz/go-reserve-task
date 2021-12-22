package domain

import "context"

type Book struct {
	ID      int64    `json:"id" binding:"required"`
	Title   string   `json:"title" binding:"required"`
	Authors []string `json:"authors" binding:"required"`
	Year    string   `json:"year" binding:"required"`
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

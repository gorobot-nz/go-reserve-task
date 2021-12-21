package book

import (
	"go-tech-task/internal/domain"

	"context"
)

type Repository interface {
	GetBooks(ctx context.Context) ([]domain.Book, error)
	GetBookById(ctx context.Context, id int64) (domain.Book, error)
	AddBooks(ctx context.Context, book domain.Book) int64
	DeleteBook(ctx context.Context, id int64) (int64, error)
	UpdateBook(ctx context.Context, id int64, book domain.Book) (int64, error)
}
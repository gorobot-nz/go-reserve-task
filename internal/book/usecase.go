package book

import (
	"context"
	"go-tech-task/internal/domain"
)

type UseCase interface {
	GetBooks(ctx context.Context) ([]domain.Book, error)
	GetBookById(ctx context.Context, id int64) (domain.Book, error)
	AddBooks(ctx context.Context, book domain.Book) (int64, error)
	DeleteBook(ctx context.Context, id int64) (int64, error)
	UpdateBook(ctx context.Context, id int64, book domain.Book) (int64, error)
}

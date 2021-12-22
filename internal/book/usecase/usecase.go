package usecase

import (
	"context"
	"go-tech-task/internal/book"
	"go-tech-task/internal/domain"
)

type BookUseCase struct {
	repository book.Repository
}

func NewBookUseCase(bookRepo book.Repository) *BookUseCase {
	return &BookUseCase{repository: bookRepo}
}

func (b *BookUseCase) GetBooks(ctx context.Context) ([]domain.Book, error) {
	return b.repository.GetBooks(ctx)
}

func (b *BookUseCase) GetBookById(ctx context.Context, id int64) (domain.Book, error) {
	return b.repository.GetBookById(ctx, id)
}

func (b *BookUseCase) AddBooks(ctx context.Context, book domain.Book) (int64, error) {
	return b.repository.AddBooks(ctx, book)
}

func (b *BookUseCase) DeleteBook(ctx context.Context, id int64) (int64, error) {
	return b.repository.DeleteBook(ctx, id)
}

func (b *BookUseCase) UpdateBook(ctx context.Context, id int64, book domain.Book) (int64, error) {
	return b.repository.UpdateBook(ctx, id, book)
}

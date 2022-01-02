package usecase

import (
	"go-tech-task/internal/domain"

	"context"
)

type BookUseCase struct {
	repository domain.BookRepository
}

func NewBookUseCase(bookRepo domain.BookRepository) *BookUseCase {
	return &BookUseCase{repository: bookRepo}
}

func (b *BookUseCase) GetBooks(ctx context.Context) ([]domain.Book, error) {
	return b.repository.GetBooks(ctx)
}

func (b *BookUseCase) GetBookById(ctx context.Context, id int64) (*domain.Book, error) {
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

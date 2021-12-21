package usecase

import (
	"go-tech-task/internal/book"
	"go-tech-task/internal/domain"
)

type BookUseCase struct {
	repository book.Repository
}

func NewBookUseCase(bookRepo book.Repository) *BookUseCase {
	return &BookUseCase{repository: bookRepo}
}

func (b *BookUseCase) GetBooks() ([]domain.Book, error) {
	return b.repository.GetBooks()
}

func (b *BookUseCase) GetBookById(id int64) (domain.Book, error) {
	return b.repository.GetBookById(id)
}

func (b *BookUseCase) AddBooks(book domain.Book) int64 {
	return b.repository.AddBooks(book)
}

func (b *BookUseCase) DeleteBook(id int64) (int64, error) {
	return b.repository.DeleteBook(id)
}

func (b *BookUseCase) UpdateBook(id int64, book domain.Book) (int64, error) {
	return b.repository.UpdateBook(id, book)
}

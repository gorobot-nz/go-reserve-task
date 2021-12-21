package usecase

import (
	"go-tech-task/internal/domain"
	"go-tech-task/internal/domain/repository"
)

type BookUsecase struct {
	repository repository.Book
}

func NewBookUsecase(repository repository.Book) *BookUsecase {
	return &BookUsecase{repository: repository}
}

func (b *BookUsecase) GetBooks() ([]domain.Book, error) {
	return b.repository.GetBooks()
}

func (b *BookUsecase) GetBookById(id int64) (domain.Book, error) {
	return b.repository.GetBookById(id)
}

func (b *BookUsecase) AddBooks(book domain.Book) int64 {
	return b.repository.AddBooks(book)
}

func (b *BookUsecase) DeleteBook(id int64) (int64, error) {
	return b.repository.DeleteBook(id)
}

func (b *BookUsecase) UpdateBook(id int64, book domain.Book) (int64, error) {
	return b.repository.UpdateBook(id, book)
}
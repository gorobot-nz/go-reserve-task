package repository

import "go-tech-task/internal/domain"

type LocalBook struct {
	books []domain.Book
}

func NewLocalBook(books []domain.Book) *LocalBook {
	return &LocalBook{books: books}
}

func (l LocalBook) GetBooks() ([]domain.Book, error) {
	panic("implement me")
}

func (l LocalBook) GetBookById(id int64) (domain.Book, error) {
	panic("implement me")
}

func (l LocalBook) AddBooks(book *domain.Book) int64 {
	panic("implement me")
}

func (l LocalBook) DeleteBook(id int64) (int64, error) {
	panic("implement me")
}

func (l LocalBook) UpdateBook(id int64) (int64, error) {
	panic("implement me")
}

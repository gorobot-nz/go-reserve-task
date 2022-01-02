package mock

import (
	"github.com/stretchr/testify/mock"

	"go-tech-task/internal/domain"

	"context"
)

type BooksPostgresStorageMock struct {
	mock.Mock
}

func (m *BooksPostgresStorageMock) GetBooks(ctx context.Context) ([]domain.Book, error) {
	args := m.Called()
	return args.Get(0).([]domain.Book), args.Error(1)
}

func (m *BooksPostgresStorageMock) GetBookById(ctx context.Context, id int64) (*domain.Book, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Book), args.Error(1)
}

func (m *BooksPostgresStorageMock) AddBooks(ctx context.Context, book domain.Book) (int64, error) {
	args := m.Called(book)
	return args.Get(0).(int64), args.Error(1)
}

func (m *BooksPostgresStorageMock) DeleteBook(ctx context.Context, id int64) (int64, error) {
	args := m.Called(id)
	return args.Get(0).(int64), args.Error(1)
}

func (m *BooksPostgresStorageMock) UpdateBook(ctx context.Context, id int64, book domain.Book) (int64, error) {
	args := m.Called(id, book)
	return args.Get(0).(int64), args.Error(1)
}

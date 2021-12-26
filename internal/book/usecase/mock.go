package usecase

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go-tech-task/internal/domain"
)

type BookUseCaseMock struct {
	mock.Mock
}

func (m BookUseCaseMock) GetBooks(ctx context.Context) ([]domain.Book, error) {
	args := m.Called()
	return args.Get(0).([]domain.Book), args.Error(1)
}

func (m BookUseCaseMock) GetBookById(ctx context.Context, id int64) (*domain.Book, error) {
	args := m.Called()
	return args.Get(0).(*domain.Book), args.Error(1)
}

func (m BookUseCaseMock) AddBooks(ctx context.Context, book domain.Book) (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func (m BookUseCaseMock) DeleteBook(ctx context.Context, id int64) (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func (m BookUseCaseMock) UpdateBook(ctx context.Context, id int64, book domain.Book) (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

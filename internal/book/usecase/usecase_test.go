package usecase

import (
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"go-tech-task/internal/book/repository/mock"
	"go-tech-task/internal/domain"

	"context"
	"testing"
)

func TestBookUseCase_AddBooks(t *testing.T) {
	rp := new(mock.BooksPostgresStorageMock)

	uc := NewBookUseCase(rp)

	testBook := domain.Book{
		Title:   "Some title",
		Authors: pq.StringArray{"First", "second"},
		Year:    "2006-01-02",
	}
	ctx := context.Background()
	rp.On("AddBooks", testBook).Return(testBook.ID, nil)
	_, err := uc.AddBooks(ctx, testBook)
	assert.NoError(t, err)
}

func TestBookUseCase_DeleteBook(t *testing.T) {
	rp := new(mock.BooksPostgresStorageMock)

	uc := NewBookUseCase(rp)

	ctx := context.Background()
	rp.On("DeleteBook", int64(1)).Return(int64(1), nil)
	id, err := uc.DeleteBook(ctx, int64(1))
	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)
}

func TestBookUseCase_GetBooks(t *testing.T) {
	books := []domain.Book{
		{
			ID:      1,
			Title:   "Check",
			Authors: pq.StringArray{"Mr Bean"},
			Year:    "1999-07-25T00:00:00Z",
		},
		{
			ID:      2,
			Title:   "Check",
			Authors: pq.StringArray{"Mr Bean"},
			Year:    "1999-07-25T00:00:00Z",
		},
	}
	rp := new(mock.BooksPostgresStorageMock)

	uc := NewBookUseCase(rp)

	ctx := context.Background()
	rp.On("GetBooks").Return(books, nil)
	result, _ := uc.GetBooks(ctx)
	assert.Equal(t, books, result)
}

func TestBookUseCase_UpdateBook(t *testing.T) {
	rp := new(mock.BooksPostgresStorageMock)

	uc := NewBookUseCase(rp)

	testBook := domain.Book{
		Title:   "Some title",
		Authors: pq.StringArray{"First", "second"},
		Year:    "2006-01-02",
	}
	ctx := context.Background()
	rp.On("UpdateBook", int64(1), testBook).Return(int64(1), nil)
	result, err := uc.UpdateBook(ctx, int64(1), testBook)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), result)
}

func TestBookUseCase_GetBookById(t *testing.T) {
	book := domain.Book{
		ID:      1,
		Title:   "Check",
		Authors: pq.StringArray{"Mr Bean"},
		Year:    "1999-07-25T00:00:00Z",
	}
	rp := new(mock.BooksPostgresStorageMock)
	uc := NewBookUseCase(rp)
	ctx := context.Background()
	rp.On("GetBookById", int64(1)).Return(&book, nil)
	result, err := uc.GetBookById(ctx, int64(1))
	assert.NoError(t, err)
	assert.Equal(t, &book, result)

}

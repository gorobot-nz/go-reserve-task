package http

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"go-tech-task/internal/book/usecase"
	"go-tech-task/internal/domain"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_AddBooks(t *testing.T) {
	book := domain.Book{
		ID:      1,
		Title:   "SomeTitle",
		Authors: pq.StringArray{"First", "second"},
		Year:    "2006-01-02",
	}

	expected, err := json.Marshal(gin.H{
		"bookId": book.ID,
	})
	assert.NoError(t, err)
	body, err := json.Marshal(book)
	assert.NoError(t, err)
	uc := new(usecase.BookUseCaseMock)
	r := gin.Default()
	RegisterEndpoints(r, uc)

	uc.On("AddBooks", book).Return(book.ID, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books/", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	actual := w.Body.Bytes()
	assert.Equal(t, string(expected), string(actual))
}

func TestHandler_GetBooks(t *testing.T) {
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
	expected, err := json.Marshal(gin.H{
		"books": books,
	})
	assert.NoError(t, err)
	uc := new(usecase.BookUseCaseMock)
	r := gin.Default()
	RegisterEndpoints(r, uc)

	uc.On("GetBooks").Return(books, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/", nil)
	r.ServeHTTP(w, req)
	actual := w.Body.Bytes()
	assert.JSONEq(t, string(expected), string(actual))
}

func TestHandler_GetBookById(t *testing.T) {
	book := &domain.Book{
		ID:      1,
		Title:   "Check",
		Authors: pq.StringArray{"Brother Koen", "Another Brother Koen"},
		Year:    "1999-07-25T00:00:00Z",
	}
	expected, err := json.Marshal(gin.H{
		"book": *book,
	})
	assert.NoError(t, err)
	uc := new(usecase.BookUseCaseMock)
	r := gin.Default()
	RegisterEndpoints(r, uc)
	uc.On("GetBookById", book.ID).Return(book, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/1", nil)
	r.ServeHTTP(w, req)
	actual := w.Body.Bytes()
	assert.JSONEq(t, string(expected), string(actual))
}

func TestHandler_DeleteBook(t *testing.T) {
	book := domain.Book{
		ID:      1,
		Title:   "SomeTitle",
		Authors: pq.StringArray{"First", "second"},
		Year:    "2006-01-02",
	}
	expected, err := json.Marshal(gin.H{
		"bookId": book.ID,
	})
	assert.NoError(t, err)

	uc := new(usecase.BookUseCaseMock)
	r := gin.Default()
	RegisterEndpoints(r, uc)

	uc.On("DeleteBook", book.ID).Return(book.ID, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/books/1", nil)
	r.ServeHTTP(w, req)
	actual := w.Body.Bytes()
	assert.Equal(t, string(expected), string(actual))
}

func TestHandler_UpdateBook(t *testing.T) {
	testBook := domain.Book{
		Title:   "SomeTitle",
		Authors: pq.StringArray{"First", "second"},
		Year:    "2006-01-02",
	}

	uc := new(usecase.BookUseCaseMock)
	r := gin.Default()
	RegisterEndpoints(r, uc)

	body, err := json.Marshal(testBook)
	assert.NoError(t, err)

	var check bytes.Buffer

	json.Unmarshal([]byte(`{"bookId":1}`), &check)

	uc.On("UpdateBook", int64(1), testBook).Return(testBook.ID, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/books/1", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	assert.Equal(t, &check, w.Body)
}

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"go-tech-task/internal/book/usecase"
	"go-tech-task/internal/domain"

	"bytes"
	"encoding/json"
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
	api := r.Group("/api")
	RegisterEndpoints(api, uc)

	uc.On("AddBooks", book).Return(book.ID, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/books", bytes.NewBuffer(body))
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
	api := r.Group("/api")
	RegisterEndpoints(api, uc)

	uc.On("GetBooks").Return(books, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/books", nil)
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
	api := r.Group("/api")
	RegisterEndpoints(api, uc)
	uc.On("GetBookById", book.ID).Return(book, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/books/1", nil)
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
	api := r.Group("/api")
	RegisterEndpoints(api, uc)

	uc.On("DeleteBook", book.ID).Return(book.ID, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/books/1", nil)
	r.ServeHTTP(w, req)
	actual := w.Body.Bytes()
	assert.Equal(t, string(expected), string(actual))
}

func TestHandler_UpdateBook(t *testing.T) {
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
	api := r.Group("/api")
	RegisterEndpoints(api, uc)

	uc.On("UpdateBook", book.ID, book).Return(book.ID, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/books/1", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	actual := w.Body.Bytes()
	assert.Equal(t, string(expected), string(actual))
}

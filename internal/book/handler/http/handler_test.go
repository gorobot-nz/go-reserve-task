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

	uc.On("AddBooks", testBook).Return(testBook.ID, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	assert.Equal(t, &check, w.Body)
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

	uc := new(usecase.BookUseCaseMock)
	r := gin.Default()
	RegisterEndpoints(r, uc)

	booksJson, err := json.Marshal(books)
	assert.NoError(t, err)

	var check bytes.Buffer
	json.Unmarshal(booksJson, &check)

	uc.On("GetBooks").Return(books, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, &check, w.Body)
}

func TestHandler_GetBookById(t *testing.T) {

}

func TestHandler_DeleteBook(t *testing.T) {
	uc := new(usecase.BookUseCaseMock)
	r := gin.Default()
	RegisterEndpoints(r, uc)

	var check *bytes.Buffer
	json.Unmarshal([]byte(`{"bookId":1}`), &check)

	uc.On("DeleteBook", int64(1)).Return(1, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/books/1", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, check, w.Body)
}

func TestHandler_UpdateBook(t *testing.T) {

}

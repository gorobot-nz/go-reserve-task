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

	var check *bytes.Buffer

	json.Unmarshal([]byte(`{"bookId":1}`), &check)

	uc.On("AddBooks", nil, testBook).Return(testBook.ID, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	assert.Equal(t, check, w.Body)
}

func TestHandler_GetBooks(t *testing.T) {

}

func TestHandler_GetBookById(t *testing.T) {

}

func TestHandler_DeleteBook(t *testing.T) {

}

func TestHandler_UpdateBook(t *testing.T) {

}

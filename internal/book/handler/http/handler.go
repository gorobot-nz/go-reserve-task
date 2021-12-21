package http

import (
	"github.com/gin-gonic/gin"

	"go-tech-task/internal/book"
	"go-tech-task/internal/domain"

	"net/http"
	"strconv"
)

type Handler struct {
	useCase book.UseCase
}

func NewHandler(useCase book.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) GetBooks(context *gin.Context) {
	books, err := h.useCase.GetBooks()

	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"books": books,
	})
}

func (h *Handler) GetBookById(context *gin.Context) {
	bookId, _ := strconv.ParseInt(context.Param("id"), 0, 64)
	book, err := h.useCase.GetBookById(bookId)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"book": book,
	})
}

func (h *Handler) AddBooks(context *gin.Context) {
	var book domain.Book

	if err := context.BindJSON(&book); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := h.useCase.AddBooks(book)
	context.JSON(http.StatusOK, map[string]interface{}{
		"bookId": id,
	})
}

func (h *Handler) DeleteBook(context *gin.Context) {
	bookId, _ := strconv.ParseInt(context.Param("id"), 0, 64)
	id, err := h.useCase.DeleteBook(bookId)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"bookId": id,
	})
}

func (h *Handler) UpdateBook(context *gin.Context) {
	bookId, _ := strconv.ParseInt(context.Param("id"), 0, 64)

	var book domain.Book

	if err := context.BindJSON(&book); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.useCase.UpdateBook(bookId, book)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"bookId": id,
	})
}

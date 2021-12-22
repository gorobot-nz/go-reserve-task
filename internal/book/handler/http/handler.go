package http

import (
	"github.com/gin-gonic/gin"
	"time"

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
	books, err := h.useCase.GetBooks(context.Request.Context())

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
	b, err := h.useCase.GetBookById(context.Request.Context(), bookId)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"book": b,
	})
}

func (h *Handler) AddBooks(context *gin.Context) {
	var b domain.Book

	if err := context.BindJSON(&b); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	num, err := strconv.ParseInt(b.Year, 0, 64)

	if err != nil || (num < 0 || num > int64(time.Now().Year())) {
		context.JSON(http.StatusBadRequest, "Wrong year format")
		return
	}

	id := h.useCase.AddBooks(context.Request.Context(), b)
	context.JSON(http.StatusOK, map[string]interface{}{
		"bookId": id,
	})
}

func (h *Handler) DeleteBook(context *gin.Context) {
	bookId, _ := strconv.ParseInt(context.Param("id"), 0, 64)
	id, err := h.useCase.DeleteBook(context.Request.Context(), bookId)

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

	var b domain.Book

	if err := context.BindJSON(&b); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.useCase.UpdateBook(context.Request.Context(), bookId, b)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"bookId": id,
	})
}

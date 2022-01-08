package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"

	"go-tech-task/internal/domain"
	"go-tech-task/pkg/middleware"

	"net/http"
)

type Handler struct {
	useCase domain.BookUseCase
}

func NewHandler(useCase domain.BookUseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) GetBooks(context *gin.Context) {
	title := context.DefaultQuery("title", "")
	books, err := h.useCase.GetBooks(context.Request.Context(), title)

	if err != nil {
		context.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"books": books,
	})
}

func (h *Handler) GetBookById(context *gin.Context) {
	bookId := context.Param("id")
	b, err := h.useCase.GetBookById(context.Request.Context(), bookId)

	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	middleware.BOOK_RESERVED.With(prometheus.Labels{
		"book_id":     bookId,
		"status_code": string(rune(http.StatusOK)),
	}).Inc()

	context.JSON(http.StatusOK, map[string]interface{}{
		"book": &b,
	})
}

func (h *Handler) AddBooks(context *gin.Context) {
	var b domain.Book

	if err := context.BindJSON(&b); err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	id, err := h.useCase.AddBooks(context.Request.Context(), b)

	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"bookId": id,
	})
}

func (h *Handler) DeleteBook(context *gin.Context) {
	bookId := context.Param("id")
	id, err := h.useCase.DeleteBook(context.Request.Context(), bookId)

	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"bookId": id,
	})
}

func (h *Handler) UpdateBook(context *gin.Context) {
	bookId := context.Param("id")

	var b domain.Book

	if err := context.BindJSON(&b); err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	id, err := h.useCase.UpdateBook(context.Request.Context(), bookId, b)

	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"bookId": id,
	})
}

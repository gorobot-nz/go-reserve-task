package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetBooks(context *gin.Context) {
	books, err := h.usecase.Book.GetBooks()

	if err != nil{
		context.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string] interface{}{
		"books": books,
	})
}

func (h *Handler) GetBookById(context *gin.Context) {
	bookId, _ := strconv.ParseInt(context.Param("id"), 0, 64)
	book, err := h.usecase.Book.GetBookById(bookId)

	if err != nil{
		context.JSON(http.StatusBadRequest,err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string] interface{}{
		"book": book,
	})
}

func (h *Handler) AddBooks(context *gin.Context) {

}

func (h *Handler) DeleteBook(context *gin.Context) {

}

func (h *Handler) UpdateBook(context *gin.Context) {

}
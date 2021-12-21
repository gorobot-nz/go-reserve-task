package http

import (
	"github.com/gin-gonic/gin"

	"go-tech-task/internal/book"
)

func RegisterEndpoints(router *gin.Engine, useCase book.UseCase){
	h := NewHandler(useCase)
	books := router.Group("/books")
	{
		books.GET("/", h.GetBooks)
		books.GET("/:id", h.GetBookById)
		books.POST("/", h.AddBooks)
		books.DELETE("/:id", h.DeleteBook)
		books.PUT("/:id", h.UpdateBook)
	}
}

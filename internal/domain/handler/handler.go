package handler

import (
	"github.com/gin-gonic/gin"
	"go-tech-task/internal/domain/usecase"
)

type Handler struct {
	usecase *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler)InitHandler() *gin.Engine{
	router := gin.New()

	api := router.Group("/api")
	{
		books := api.Group("/books")
		{
			books.GET("/", h.GetBooks)
			books.GET("/:id", h.GetBookById)
			books.POST("/", h.AddBooks)
			books.DELETE("/:id", h.DeleteBook)
			books.PUT("/:id", h.UpdateBook)
		}
	}
	return router
}
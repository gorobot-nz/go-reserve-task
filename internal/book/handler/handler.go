package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	
}

func NewHandler() *Handler {
	return &Handler{}
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
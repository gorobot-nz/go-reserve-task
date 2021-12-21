package server

import (
	"github.com/gin-gonic/gin"

	"go-tech-task/internal/book"
	bookHTTP "go-tech-task/internal/book/handler/http"
	localDB "go-tech-task/internal/book/repository/local"
	bookUseCase "go-tech-task/internal/book/usecase"
	"go-tech-task/internal/domain"

	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	server *http.Server

	bookUC book.UseCase
}

func NewApp() *App {
	localDb := []domain.Book{
		{ID: 1, Title: "Fight Club", Authors: []string{"Palahniuc"}, Year: "1996"},
		{ID: 2, Title: "Theoretical Physics Course", Authors: []string{"Landau", "Lifshitz"}, Year: "1954"},
		{ID: 3, Title: "Reptiloids", Authors: []string{"Prokopenko"}, Year: "2015"},
		{ID: 4, Title: "Another reptiloids", Authors: []string{"Prokopenko, Chapman"}, Year: "2017"},
		{ID: 5, Title: "Once upon a time in Hollywood", Authors: []string{"Tarantino"}, Year: "2019"},
		{ID: 6, Title: "Computer architecture", Authors: []string{"Tanenbaum"}, Year: "1975"},
		{ID: 7, Title: "Making a compact hydrogen bomb in labor lessons", Authors: []string{"Makarenko"}, Year: "1960"},
		{ID: 8, Title: "Code: The Hidden Language of Computer Hardware and Software", Authors: []string{"Petzold"}, Year: "1999"},
	}

	bookRepo := localDB.NewBooksLocalStorage(localDb)

	return &App{
		bookUC: bookUseCase.NewBookUseCase(bookRepo),
	}
}

func (a *App) Run() error {

	router := gin.Default()

	bookHTTP.RegisterEndpoints(router, a.bookUC)

	a.server = &http.Server{
		Addr:           ":8000",
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)

	defer shutdown()
	return a.server.Shutdown(ctx)
}

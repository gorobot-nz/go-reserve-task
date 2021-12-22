package server

import (
	"github.com/gin-gonic/gin"

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

	bookUC domain.BookUseCase
}

func NewApp() *App {
	localDb := []domain.Book{
		{ID: 1, Title: "Fight Club", Authors: []string{"Palahniuc"}, Year: time.Date(2006, 1, 2, 15, 04, 05, 0, time.UTC)},
		{ID: 2, Title: "Theoretical Physics Course", Authors: []string{"Landau", "Lifshitz"}, Year: time.Date(2008, 1, 2, 15, 04, 05, 0, time.UTC)},
		{ID: 3, Title: "Reptiloids", Authors: []string{"Prokopenko"}, Year: time.Date(1995, 1, 2, 15, 04, 05, 0, time.UTC)},
		{ID: 4, Title: "Another reptiloids", Authors: []string{"Prokopenko, Chapman"}, Year: time.Date(1998, 1, 2, 15, 04, 05, 0, time.UTC)},
		{ID: 5, Title: "Once upon a time in Hollywood", Authors: []string{"Tarantino"}, Year: time.Date(2019, 1, 2, 15, 04, 05, 0, time.UTC)},
		{ID: 6, Title: "Computer architecture", Authors: []string{"Tanenbaum"}, Year: time.Date(2020, 1, 2, 15, 04, 05, 0, time.UTC)},
		{ID: 7, Title: "Making a compact hydrogen bomb in labor lessons", Authors: []string{"Makarenko"}, Year: time.Date(2001, 1, 2, 15, 04, 05, 0, time.UTC)},
		{ID: 8, Title: "Code: The Hidden Language of Computer Hardware and Software", Authors: []string{"Petzold"}, Year: time.Date(2003, 1, 2, 15, 04, 05, 0, time.UTC)},
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

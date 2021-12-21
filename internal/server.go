package internal

import (
	"context"
	"go-tech-task/internal/domain"
	"go-tech-task/internal/handler"
	"go-tech-task/internal/repository"
	"go-tech-task/internal/usecase"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (s *Server) Run() error {
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

	repos := repository.NewRepository(localDb)
	ucase := usecase.NewUsecase(repos)
	h := handler.NewHandler(ucase)

	s.server = &http.Server{
		Addr:           ":8000",
		Handler: h.InitHandler(),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

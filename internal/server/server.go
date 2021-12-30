package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-tech-task/pkg/middleware"

	bookHTTP "go-tech-task/internal/book/handler/http"
	"go-tech-task/internal/book/repository/postgres"
	bookUseCase "go-tech-task/internal/book/usecase"
	"go-tech-task/internal/domain"

	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func checkEnvVars() {
	requiredEnvs := []string{"POSTGRES_HOST", "POSTGRES_USER", "POSTGRES_PASSWORD"}
	var msg []string
	for _, el := range requiredEnvs {
		val, exists := os.LookupEnv(el)
		if !exists || len(val) == 0 {
			msg = append(msg, el)
		}
	}
	if len(msg) > 0 {
		log.Fatal(strings.Join(msg, ", "), " env(s) not set")
	}
}

type App struct {
	server *http.Server

	bookUC domain.BookUseCase
}

func NewApp() *App {
	if err := InitConfig(); err != nil {
		log.Fatalf("Config error: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Env error: %s", err.Error())
	}

	checkEnvVars()

	config := postgres.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     viper.GetString("db.POSTGRES_DBPORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   viper.GetString("db.POSTGRES_DBNAME"),
		SSLMode:  viper.GetString("db.POSTGRES_SSLMODE"),
	}

	bookRepo := postgres.NewBooksPostgresStorage(config)

	return &App{
		bookUC: bookUseCase.NewBookUseCase(bookRepo),
	}
}

func (a *App) Run() error {

	router := gin.Default()
	router.GET("/metrics", func(c *gin.Context) {
		handler := promhttp.Handler()
		handler.ServeHTTP(c.Writer, c.Request)
	})
	api := router.Group("/api")

	bookHTTP.RegisterEndpoints(api, a.bookUC)

	metricsMw := middleware.NewPrometheusMiddleware("books")

	router.Use(metricsMw.Metrics())
	router.Use(middleware.CORS())
	router.Use(middleware.Logging())

	a.server = &http.Server{
		Addr:           ":" + viper.GetString("port"),
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

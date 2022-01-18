package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "go-tech-task/docs"

	"go-tech-task/internal/book/repository/elastic_book"

	bookHTTP "go-tech-task/internal/book/handler/http"
	bookUseCase "go-tech-task/internal/book/usecase"
	"go-tech-task/internal/domain"
	"go-tech-task/pkg/middleware"

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
	requiredEnvs := []string{"ELASTIC_HOST", "ELASTIC_USER", "ELASTIC_PASSWORD"}
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

func logInit() {
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
}

type App struct {
	server *http.Server

	bookUC domain.BookUseCase
}

func NewApp() *App {
	logInit()
	if err := InitConfig(); err != nil {
		log.Fatalf("Config error: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Env error: %s", err.Error())
	}

	checkEnvVars()

	config := elastic_book.Config{
		Host:     os.Getenv("ELASTIC_HOST"),
		Username: os.Getenv("ELASTIC_USER"),
		Password: os.Getenv("ELASTIC_PASSWORD"),
		Index:    viper.GetString("index"),
	}

	bookRepo := elastic_book.NewBooksElasticStorage(config)

	return &App{
		bookUC: bookUseCase.NewBookUseCase(bookRepo),
	}
}

func (a *App) Run() error {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/metrics", func(c *gin.Context) {
		handler := promhttp.Handler()
		handler.ServeHTTP(c.Writer, c.Request)
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	metricsMw := middleware.NewPrometheusMiddleware("books")

	api := router.Group("/api")
	api.Use(metricsMw.Metrics())
	api.Use(middleware.CORS())
	api.Use(middleware.Logging())

	bookHTTP.RegisterEndpoints(api, a.bookUC)
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

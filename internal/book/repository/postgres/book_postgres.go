package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-tech-task/internal/domain"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const layout = "2006-01-02"

type BooksPostgresStorage struct {
	conn *sqlx.DB
}

func NewBooksPostgresStorage(cfg Config) *BooksPostgresStorage {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		log.Fatalf("DBConnection error: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("DBConnection error: %s", err.Error())
	}

	path := filepath.Join(".", "internal", "schema", "book.sql")

	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		log.Fatalf("DBConnection error: %s", err.Error())
	}

	var schema = string(c)
	db.MustExec(schema)

	return &BooksPostgresStorage{conn: db}
}

func (b *BooksPostgresStorage) GetBooks(ctx context.Context) ([]domain.Book, error) {
	var books []domain.Book
	query := fmt.Sprintf("SELECT id, title, authors, book_year FROM %s", "books")
	err := b.conn.Select(&books, query)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return books, nil
}

func (b *BooksPostgresStorage) GetBookById(ctx context.Context, id int64) (domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BooksPostgresStorage) AddBooks(ctx context.Context, book domain.Book) (int64, error) {
	var id int64
	date, _ := time.Parse(layout, book.Year)
	query := fmt.Sprintf("INSERT INTO %s (title, authors, book_year) values ($1, $2, $3) RETURNING id", "books")
	row := b.conn.QueryRow(query, book.Title, book.Authors, date)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (b *BooksPostgresStorage) DeleteBook(ctx context.Context, id int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BooksPostgresStorage) UpdateBook(ctx context.Context, id int64, book domain.Book) (int64, error) {
	//TODO implement me
	panic("implement me")
}

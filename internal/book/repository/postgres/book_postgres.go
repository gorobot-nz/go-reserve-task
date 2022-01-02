package postgres

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	"go-tech-task/internal/domain"

	"context"
	"errors"
	"fmt"
	"io/ioutil"
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
		log.Fatalf("DBConnection error: %s", ioErr.Error())
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

func (b *BooksPostgresStorage) GetBookById(ctx context.Context, id int64) (*domain.Book, error) {
	var book domain.Book
	query := fmt.Sprintf("SELECT id, title, authors, book_year FROM %s WHERE id = $1", "books")
	err := b.conn.Get(&book, query, id)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (b *BooksPostgresStorage) AddBooks(ctx context.Context, book domain.Book) (int64, error) {
	var id int64
	date, err := time.Parse(layout, book.Year)

	if err != nil {
		return 0, err
	}

	err = Validation(&book)

	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (title, authors, book_year) values ($1, $2, $3) RETURNING id", "books")
	row := b.conn.QueryRow(query, book.Title, book.Authors, date)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (b *BooksPostgresStorage) DeleteBook(ctx context.Context, id int64) (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", "books")
	_, err := b.conn.Exec(query, id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (b *BooksPostgresStorage) UpdateBook(ctx context.Context, id int64, book domain.Book) (int64, error) {
	date, err := time.Parse(layout, book.Year)

	if err != nil {
		return 0, err
	}

	err = Validation(&book)

	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("UPDATE %s SET title = $1, authors = $2, book_year = $3, updated_at = $4 WHERE id = $5", "books")
	_, err = b.conn.Exec(query, book.Title, book.Authors, date, time.Now(), id)

	if err != nil {
		return 0, nil
	}

	return id, err
}

func Validation(book *domain.Book) error {
	if len(book.Authors) == 0 {
		return errors.New("Wrong author format")
	}

	for _, value := range book.Authors {
		if value == "" {
			return errors.New("Wrong author format")
		}
	}

	return nil
}

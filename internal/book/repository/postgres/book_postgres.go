package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"go-tech-task/internal/domain"
	"log"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

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

	var schema = `
	CREATE TABLE IF NOT EXISTS books(
		id serial PRIMARY KEY NOT NULL UNIQUE,
		title varchar(255) NOT NULL,
		authors varchar(255)[],
		book_year timestamp 
	)
	`
	db.MustExec(schema)

	return &BooksPostgresStorage{conn: db}
}

func (b *BooksPostgresStorage) GetBooks(ctx context.Context) ([]domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BooksPostgresStorage) GetBookById(ctx context.Context, id int64) (domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BooksPostgresStorage) AddBooks(ctx context.Context, book domain.Book) (int64, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (title, authors, book_year) values ($1, $2, $3) RETURNING id", "books")
	row := b.conn.QueryRow(query, book.Title, pq.Array(book.Authors), pq.FormatTimestamp(book.Year))
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

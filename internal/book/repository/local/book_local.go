package local

import (
	"context"
	"go-tech-task/internal/domain"
	"strconv"
	"time"

	"errors"
)

type BooksLocalStorage struct {
	books []domain.Book
}

func NewBooksLocalStorage(books []domain.Book) *BooksLocalStorage {
	return &BooksLocalStorage{books: books}
}

func (l *BooksLocalStorage) GetBooks(ctx context.Context) ([]domain.Book, error) {
	return l.books, nil
}

func (l *BooksLocalStorage) GetBookById(ctx context.Context, id int64) (domain.Book, error) {
	for _, value := range l.books {
		if value.ID == id {
			return value, nil
		}
	}
	return domain.Book{}, errors.New("No such id")
}

func (l *BooksLocalStorage) AddBooks(ctx context.Context, book domain.Book) (int64, error) {

	if len(book.Authors) == 0 {
		return 0, errors.New("Wrong author format")
	}

	for _, value := range book.Authors{
		if value == "" {
			return 0, errors.New("Wrong author format")
		}
	}

	num, err := strconv.ParseInt(book.Year, 0, 64)

	if err != nil || (num < 0 || num > int64(time.Now().Year())) {
		return 0, errors.New("Wrong year format")
	}

	l.books = append(l.books, book)
	return book.ID, nil
}

func (l *BooksLocalStorage) DeleteBook(ctx context.Context, id int64) (int64, error) {
	var bookId int64 = -1
	var bookIndex int64 = -1

	for index, value := range l.books {
		if value.ID == id {
			bookIndex = int64(index)
			bookId = value.ID
			break
		}
	}

	if bookId < 0 {
		return bookId, errors.New("No such id")
	}

	l.books = append(l.books[:bookIndex], l.books[bookIndex+1:]...)
	return bookId, nil
}

func (l *BooksLocalStorage) UpdateBook(ctx context.Context, id int64, book domain.Book) (int64, error) {
	for i := 0; i < len(l.books); i++ {
		if l.books[i].ID == id {
			if book.Year != "" {
				l.books[i].Year = book.Year
			}
			if len(book.Authors) != 0 {
				l.books[i].Authors = book.Authors
			}
			if book.Title != "" {
				l.books[i].Title = book.Title
			}
			return id, nil
		}
	}
	return 0, errors.New("No such id")
}

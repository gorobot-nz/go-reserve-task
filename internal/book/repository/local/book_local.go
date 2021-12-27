package local

import (
	"context"
	"errors"
	"go-tech-task/internal/domain"
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

func (l *BooksLocalStorage) GetBookById(ctx context.Context, id int64) (*domain.Book, error) {
	for _, value := range l.books {
		if value.ID == id {
			return &value, nil
		}
	}
	return nil, errors.New("No such id")
}

func (l *BooksLocalStorage) AddBooks(ctx context.Context, book domain.Book) (int64, error) {

	err := Validation(&book)

	if err != nil {
		return 0, err
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
			err := Validation(&l.books[i])
			if err != nil {
				return 0, err
			}
			l.books[i] = book
			l.books[i].ID = id
			return id, nil
		}
	}
	return 0, errors.New("No such id")
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

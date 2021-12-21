package local

import (
	"go-tech-task/internal/domain"

	"errors"
)

type BooksLocalStorage struct {
	books []domain.Book
}

func NewBooksLocalStorage(books []domain.Book) *BooksLocalStorage {
	return &BooksLocalStorage{books: books}
}

func (l *BooksLocalStorage) GetBooks() ([]domain.Book, error) {
	return l.books, nil
}

func (l *BooksLocalStorage) GetBookById(id int64) (domain.Book, error) {
	for _, value := range l.books {
		if value.ID == id {
			return value, nil
		}
	}
	return domain.Book{}, errors.New("No such id")
}

func (l *BooksLocalStorage) AddBooks(book domain.Book) int64 {
	l.books = append(l.books, book)
	return book.ID
}

func (l *BooksLocalStorage) DeleteBook(id int64) (int64, error) {
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

func (l *BooksLocalStorage) UpdateBook(id int64, book domain.Book) (int64, error) {
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

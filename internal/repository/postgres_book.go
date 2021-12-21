package repository

import (
	"errors"
	"fmt"
	"go-tech-task/internal/domain"
)

type LocalBook struct {
	books []domain.Book
}

func NewLocalBook(books []domain.Book) *LocalBook {
	return &LocalBook{books: books}
}

func (l *LocalBook) GetBooks() ([]domain.Book, error) {
	return l.books, nil
}

func (l *LocalBook) GetBookById(id int64) (domain.Book, error) {
	for _, value := range l.books{
		if value.ID == id{
			return value, nil
		}
	}
	return domain.Book{}, errors.New("No such id")
}

func (l *LocalBook) AddBooks(book domain.Book) int64 {
	fmt.Println(l.books)
	fmt.Println(book)
	l.books = append(l.books, book)
	fmt.Println(l.books)
	return book.ID
}

func (l *LocalBook) DeleteBook(id int64) (int64, error) {
	panic("implement me")
}

func (l *LocalBook) UpdateBook(id int64) (int64, error) {
	panic("implement me")
}

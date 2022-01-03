package elastic

import (
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"

	"go-tech-task/internal/domain"

	"context"
)

type BooksElasticStorage struct {
	client *elastic.Client
}

func NewBooksElasticStorage() *BooksElasticStorage {
	client, err := elastic.NewClient()
	if err != nil {
		logrus.Fatalf("Error elastic client: %+v", err)
	}
	return &BooksElasticStorage{client}
}

func (b *BooksElasticStorage) GetBooks(ctx context.Context) ([]domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BooksElasticStorage) GetBookById(ctx context.Context, id int64) (*domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BooksElasticStorage) AddBooks(ctx context.Context, book domain.Book) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BooksElasticStorage) DeleteBook(ctx context.Context, id int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BooksElasticStorage) UpdateBook(ctx context.Context, id int64, book domain.Book) (int64, error) {
	//TODO implement me
	panic("implement me")
}

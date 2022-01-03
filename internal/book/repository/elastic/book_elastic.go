package elastic

import (
	"fmt"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"

	"go-tech-task/internal/domain"

	"context"
)

type BooksElasticStorage struct {
	client *elastic.Client
}

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"book":{
			"properties":{
				"authors":{
					"type":"text"
				},
				"title":{
					"type":"keyword"
				},
				"year":{
					"type":"date"
				},
			}
		}
	}
}`

const hostDb = "http://127.0.0.1:9200"

func NewBooksElasticStorage() *BooksElasticStorage {
	ctx := context.Background()
	client, err := elastic.NewClient()
	if err != nil {
		logrus.Fatalf("Error elastic client: %+v", err)
	}

	info, code, err := client.Ping(hostDb).Do(ctx)
	if err != nil {
		logrus.Fatalf("Error ping: %+v", err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	exists, err := client.IndexExists("Books").Do(ctx)
	if err != nil {
		logrus.Fatalf("Books index check error: %+v", err)
	}
	if !exists {
		createIndex, err := client.CreateIndex("Books").BodyString(mapping).Do(ctx)
		if err != nil {
			logrus.Fatalf("Books index creating error: %+v", err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
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

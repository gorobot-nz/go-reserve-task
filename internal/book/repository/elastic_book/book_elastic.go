package elastic_book

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"go-tech-task/internal/domain"
	"strconv"
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
					"type":"text"
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
		// Handle error
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(hostDb).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion(hostDb)
	if err != nil {
		// Handle error
		logrus.Fatalf("error %+v", err.Error())
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists("books").Do(ctx)
	if err != nil {
		// Handle error
		logrus.Fatalf("error %+v", err.Error())
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("books").BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			logrus.Fatalf("error %+v", err.Error())
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
	_, err := b.client.Get().
		Index("books").
		Type("book").
		Id(strconv.FormatInt(id, 10)).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (b *BooksElasticStorage) AddBooks(ctx context.Context, book domain.Book) (int64, error) {
	put1, err := b.client.Index().
		Index("books").
		Type("book").
		Id(strconv.FormatInt(book.ID, 10)).
		BodyJson(book).
		Do(ctx)
	if err != nil {
		return 0, err
	}
	logrus.Infof("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return book.ID, nil
}

func (b *BooksElasticStorage) DeleteBook(ctx context.Context, id int64) (int64, error) {
	b.client.Delete()
}

func (b *BooksElasticStorage) UpdateBook(ctx context.Context, id int64, book domain.Book) (int64, error) {
	//TODO implement me
	panic("implement me")
}

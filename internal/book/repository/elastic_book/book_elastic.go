package elastic_book

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"go-tech-task/internal/domain"
)

type BooksElasticStorage struct {
	client *elastic.Client
}

const mapping = `
{
	"mappings": {
		"properties": {
			"authors": { "type": "keyword" },
			"title": { "type": "text" },
			"year": { "type": "date" }	
		}
  	},
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	}
}`

const hostDb = "http://127.0.0.1:9200"

type ElasticBook struct {
	Title   string   `json:"title" binding:"required" db:"title"`
	Authors []string `json:"authors" db:"authors"`
	Year    string   `json:"year" binding:"required" db:"book_year"`
}

func NewBooksElasticStorage() *BooksElasticStorage {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetBasicAuth("elastic", "chageme"))
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
		logrus.Fatalf("error %+v", err.Error())
	}

	if !exists {
		// Create a new index.
		_, err := client.CreateIndex("books").BodyString(mapping).Do(ctx)
		fmt.Println("Create")
		if err != nil {
			// Handle error
			logrus.Fatalf("error %+v", err.Error())
		}
	}
	return &BooksElasticStorage{client}
}

func (b *BooksElasticStorage) GetBooks(ctx context.Context) ([]domain.Book, error) {
	_, err := b.client.Get().
		Index("books").
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return []domain.Book{}, nil
}

func (b *BooksElasticStorage) GetBookById(ctx context.Context, id string) (*domain.Book, error) {
	_, err := b.client.Get().
		Index("books").
		Id(id).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (b *BooksElasticStorage) AddBooks(ctx context.Context, book domain.Book) (string, error) {
	put1, err := b.client.Index().
		Index("books").
		BodyJson(ElasticBook{
			Year:    book.Year,
			Authors: book.Authors,
			Title:   book.Title,
		}).
		Do(ctx)
	if err != nil {
		return "0", err
	}
	logrus.Infof("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return put1.Id, nil
}

func (b *BooksElasticStorage) DeleteBook(ctx context.Context, id string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BooksElasticStorage) UpdateBook(ctx context.Context, id string, book domain.Book) (string, error) {
	//TODO implement me
	panic("implement me")
}

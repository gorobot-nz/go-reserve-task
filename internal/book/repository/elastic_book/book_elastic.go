package elastic_book

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"go-tech-task/internal/domain"
	"reflect"
)

type BooksElasticStorage struct {
	client *elastic.Client
}

const mapping = `
{
    "settings": {
        "index": {
            "number_of_shards": 2,
            "number_of_replicas": 1,
            "analysis": {
                "analyzer": {
                    "custom_russian": {
                        "type": "custom",
                        "tokenizer": "standard",
                        "char_filter": [
                            "custom_to_russian"
                        ],
                        "filter": [
                            "lowercase"
                        ]
                    },
                    "custom_english": {
                        "type": "custom",
                        "tokenizer": "standard",
                        "char_filter": [
                            "custom_to_english"
                        ],
                        "filter": [
                            "lowercase"
                        ]
                    }
                },
                "char_filter": {
                    "custom_to_russian": {
                        "type": "mapping",
                        "mappings": [
                            "a => А",
                            "b => Б",
                            "c => К",
                            "d => Д",
                            "e => Е",
                            "f => Ф",
                            "g => Г",
                            "h => Х",
                            "i => И",
                            "j => ДЖ",
                            "k => К",
                            "l => Л",
                            "m => М",
                            "n => Н",
                            "o => О",
                            "p => П",
                            "q => К",
                            "r => Р",
                            "s => С",
                            "t => Т",
                            "u => Ю",
                            "v => В",
                            "w => В",
                            "x => КС",
                            "y => АЙ",
                            "z => З"
                        ]
                    },
                    "custom_to_english": {
                        "type": "mapping",
                        "mappings": [
                            "а => A",
                            "б => B",
                            "в => V",
                            "г => G",
                            "д => D",
                            "е => E",
                            "ё => YO",
                            "ж => ZH",
                            "з => Z",
                            "и => I",
                            "к => K",
                            "л => L",
                            "м => M",
                            "н => N",
                            "о => O",
                            "п => P",
                            "р => R",
                            "с => S",
                            "т => T",
                            "у => U",
                            "ф => F",
                            "х => H",
                            "ц => C",
                            "ч => CH",
                            "ш => SH",
                            "щ => SHCH",
                            "э => E",
                            "ю => U",
                            "я => YA"
                        ]
                    }
                }
            }
        }
    },
    "mappings": {
        "properties": {
            "authors": {
                "type": "keyword"
            },
            "title": {
                "type": "text",
                "fields": {
                    "en": {
                        "type": "text",
                        "analyzer": "custom_english"
                    },
                    "ru": {
                        "type": "text",
                        "analyzer": "custom_russian"
                    }
                }
            },
            "year": {
                "type": "date"
            }
        }
    }
}
`

const hostDb = "http://127.0.0.1:9200"

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

func (b *BooksElasticStorage) GetBooks(ctx context.Context, title string) ([]domain.Book, error) {

	query := elastic.NewMultiMatchQuery(title, "title", "title.en", "title.ru").
		Operator("and").
		Fuzziness("AUTO")

	var result *elastic.SearchResult
	var err error

	if len(title) == 0 {
		result, err = b.client.Search().
			Index("books").
			Do(ctx)
	} else {
		result, err = b.client.Search().
			Index("books").
			Query(query).
			Do(ctx)
	}

	var book domain.Book
	var books []domain.Book

	for _, item := range result.Each(reflect.TypeOf(book)) {
		t := item.(domain.Book)
		book.Year = t.Year
		book.Title = t.Title
		book.Authors = t.Authors
		books = append(books, book)
	}

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *BooksElasticStorage) GetBookById(ctx context.Context, id string) (*domain.Book, error) {
	termQuery := elastic.NewTermQuery("_id", id)

	result, err := b.client.Search().
		Index("books").
		Query(termQuery).
		Do(ctx)

	var book domain.Book

	for _, item := range result.Each(reflect.TypeOf(book)) {
		t := item.(domain.Book)
		book.Year = t.Year
		book.Title = t.Title
		book.Authors = t.Authors
	}

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (b *BooksElasticStorage) AddBooks(ctx context.Context, book domain.Book) (string, error) {
	put1, err := b.client.Index().
		Index("books").
		BodyJson(book).
		Do(ctx)
	if err != nil {
		return "0", err
	}
	return put1.Id, nil
}

func (b *BooksElasticStorage) DeleteBook(ctx context.Context, id string) (string, error) {
	res, err := b.client.Delete().Index("books").
		Id(id).Refresh("true").Do(ctx)
	if err != nil {
		return "0", err
	}
	return res.Id, nil
}

func (b *BooksElasticStorage) UpdateBook(ctx context.Context, id string, book domain.Book) (string, error) {
	res, err := b.client.Update().
		Index("books").
		Id(id).
		Doc(map[string]interface{}{
			"title":   book.Title,
			"authors": book.Authors,
			"year":    book.Year,
		}).
		Do(ctx)
	if err != nil {
		return "0", err
	}
	return res.Id, nil
}

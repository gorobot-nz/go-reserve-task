package elastic

import (
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
)

type BooksElasticStorage struct {
	client *elastic.Client
}

func NewBooksElasticStorage() *BooksElasticStorage {
	client, err := elastic.NewClient()
	if err != nil {
		logrus.Fatal("Error")
	}
	return &BooksElasticStorage{client}
}

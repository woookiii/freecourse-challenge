package repository

import (
	"log"
	"worker-search/config"

	"github.com/elastic/go-elasticsearch/v9"
)

type Repository struct {
	typedClient *elasticsearch.TypedClient
}

func NewRepository(config *config.Config) *Repository {
	r := new(Repository)

	cfg := elasticsearch.Config{
		Addresses: config.Elasticsearch.URLS,
	}
	var err error
	r.typedClient, err = elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Panicf("fail to create Elasticsearch typedClient: %v", err)
	}

	return r
}

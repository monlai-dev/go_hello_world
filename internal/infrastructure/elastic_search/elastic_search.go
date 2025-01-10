package elastic_search

import (
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
)

var elasticClient *elasticsearch8.Client

func InitElasticSearch() *elasticsearch8.Client {
	cfg := elasticsearch8.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	client, err := elasticsearch8.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	elasticClient = client
	return client
}

func GetElasticSearchClient() *elasticsearch8.Client {
	return elasticClient
}

func CreateIndex(indexName string) {
	_, err := elasticClient.Indices.Create(
		indexName,
	)

	if err != nil {
		panic(err)
	}
}

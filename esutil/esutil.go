package esutil

import (
	"context"
	"fmt"
	"search_engine/crawler"
	"strconv"

	"github.com/olivere/elastic"
)

// CreateClient provides es client
func CreateClient() (*elastic.Client, context.Context) {
	ctx := context.Background()

	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}

	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	return client, ctx
}

// CreateIndex creates index given name and mapping
func CreateIndex(ctx context.Context, es *elastic.Client, indexName string, mapping string) {
	exists, err := es.IndexExists(indexName).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		createIndex, err := es.CreateIndex(indexName).BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			fmt.Println("Not acknowledged")
		}
	}
}

// IndexQuote indexes a Quote
func IndexQuote(ctx context.Context, es *elastic.Client, indexName string, typeName string, id int, quote crawler.Quote) {
	put, err := es.Index().
		Index(indexName).
		Type(typeName).
		Id(strconv.Itoa(id)).
		BodyJson(quote).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed quote %s to index %s, type %s\n", put.Id, put.Index, put.Type)
}

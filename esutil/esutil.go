package esutil

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"search_engine/crawler"
	"strconv"

	"github.com/olivere/elastic/v7"
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
	fmt.Printf("\nElasticsearch returned with code %d and version %s\n", code, info.Version.Number)

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

// BulkIndexQuotes bulk indexes quotes crawled
func BulkIndexQuotes(ctx context.Context, es *elastic.Client, indexName string, quotes []crawler.Quote) {
	bulk := es.Bulk()
	for i, quote := range quotes {
		idStr := strconv.Itoa(i)
		req := elastic.NewBulkIndexRequest()
		req.OpType("index")
		req.Index(indexName)
		req.Id(idStr)
		req.Doc(quote)
		bulk = bulk.Add(req)
	}
	bulkResp, err := bulk.Do(ctx)

	if err != nil {
		log.Fatalf("bulk.Do(ctx) ERROR: %s", err)
	} else {
		indexed := bulkResp.Indexed()
		fmt.Println("bulkResp.Indexed():", indexed)
	}
}

// SearchIndex for our queries
func SearchIndex(ctx context.Context, es *elastic.Client, indexName string, query string) {
	var searchResult *elastic.SearchResult
	var err error

	highlight := elastic.NewHighlight().Field("content").NumOfFragments(5).FragmentSize(25)
	highlight = elastic.NewHighlight().Field("author").NumOfFragments(0)
	highlight = highlight.PreTags("<span>").PostTags("</span>")

	searchResult, err = es.Search().
		Index(indexName).
		Highlight(highlight).
		Query(elastic.RawStringQuery(query)).
		SortBy(elastic.NewFieldSort("_doc"), elastic.NewFieldSort("_score").Desc()).
		From(0).Size(50).
		Pretty(true).
		Do(ctx)

	if err != nil {
		panic(err)
	} else {
		fmt.Println(searchResult.TookInMillis)
		fmt.Println(searchResult.Hits.TotalHits)

		var quotes []crawler.Quote
		for _, hit := range searchResult.Hits.Hits {
			var quote crawler.Quote
			err := json.Unmarshal(hit.Source, &quote)
			if err != nil {
				fmt.Println("[Getting Quotes [Unmarshal] Err=", err)
			}
			quotes = append(quotes, quote)
		}
		if err != nil {
			fmt.Println("Fetching quote fail: ", err)
		} else {
			for _, q := range quotes {
				fmt.Printf("%s, %s\n", q.Content, q.Author)
			}
		}
	}

}

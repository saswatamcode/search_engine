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

//SearchRes is the search result we need
type SearchRes struct {
	Milliseconds int64           `json:"milliseconds"`
	TotalHit     int64           `json:"totalHits"`
	Quotes       []crawler.Quote `json:"quotes"`
}

var elasticSearchURL string = "http://127.0.0.1:9200"

// CreateClient provides es client
func CreateClient() (*elastic.Client, context.Context) {
	ctx := context.Background()

	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}

	info, code, err := client.Ping(elasticSearchURL).Do(ctx)
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
		idStr := strconv.Itoa(i + 1)
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
func SearchIndex(ctx context.Context, es *elastic.Client, indexName string, query string) (SearchRes, error) {
	var searchResult *elastic.SearchResult
	var err error

	var result SearchRes

	highlight := elastic.NewHighlight().Field("content").NumOfFragments(5).FragmentSize(25)
	highlight = elastic.NewHighlight().Field("author").NumOfFragments(0)
	highlight = highlight.PreTags("<span>").PostTags("</span>")

	searchResult, err = es.Search().
		Index(indexName).
		Highlight(highlight).
		Query(elastic.RawStringQuery(query)).
		SortBy(elastic.NewFieldSort("_doc"), elastic.NewFieldSort("_score").Desc()).
		From(0).Size(100).
		Pretty(true).
		Do(ctx)
	if err != nil {
		return result, err
	}
	fmt.Printf("\nSearch took %d milliseconds!", searchResult.TookInMillis)
	fmt.Printf("\nFound %d quotes!\n", searchResult.Hits.TotalHits.Value)
	result.Milliseconds = searchResult.TookInMillis
	result.TotalHit = searchResult.Hits.TotalHits.Value

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
		return result, err
	}
	if len(quotes) > 0 {
		fmt.Println("\nYour search results are: ")
		for _, q := range quotes {
			fmt.Printf("%s, %s\n", q.Content, q.Author)
		}
		result.Quotes = quotes
		return result, nil
	}
	fmt.Println("You have no matching quotes. Try indexing more quotes!")
	return result, nil
}

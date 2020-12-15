package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

func main() {
	// quotes := crawler.Run(100)
	// for _, quote := range quotes {
	// 	fmt.Println(quote.Content)
	// }
	ctx := context.Background()

	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}

	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

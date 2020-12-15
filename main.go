package main

import (
	"fmt"
	"search_engine/crawler"
)

func main() {
	quotes := crawler.Run(100)
	for _, quote := range quotes {
		fmt.Println(quote.Content)
	}
}

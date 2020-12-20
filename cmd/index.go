package cmd

import (
	"fmt"
	"search_engine/crawler"
	"search_engine/esutil"

	"github.com/spf13/cobra"
)

var (
	indexSetup bool
	maxQuotes  int
)

func init() {
	rootCmd.AddCommand(indexCmd)
	indexCmd.Flags().IntVarP(&maxQuotes, "maxQuotes", "q", 100, "Maximum number of quotes to be indexed")
	indexCmd.Flags().BoolVarP(&indexSetup, "setup", "s", true, "Create Elasticsearch index")
}

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Indexes quotes into Elasticsearch by crawling goodreads.com",
	Run: func(cmd *cobra.Command, args []string) {
		quotes := crawler.Run(maxQuotes)
		fmt.Printf("Scraped %d quotes", len(quotes))

		client, ctx := esutil.CreateClient()
		if indexSetup {
			fmt.Println("Creating Index")
			esutil.CreateIndex(ctx, client, IndexName, esutil.Mapping)
		}
		fmt.Println("Bulk indexing quotes:")
		esutil.BulkIndexQuotes(ctx, client, IndexName, quotes)
	},
}

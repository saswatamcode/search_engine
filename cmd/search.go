package cmd

import (
	"fmt"
	"search_engine/esutil"
	"strings"

	"github.com/spf13/cobra"
)

var searchAll bool

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolVarP(&searchAll, "searchall", "a", false, "Query all data from Elasticsearch index")
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches quotes from goodreads.com",
	Run: func(cmd *cobra.Command, args []string) {
		client, ctx := esutil.CreateClient()

		var query string
		if searchAll {
			query = esutil.QueryAll
		} else {
			if len(args) < 1 {
				fmt.Println("You need to provide a search string!")
				return
			}
			query = fmt.Sprintf(esutil.SearchQuery, strings.Join(args, " "))
		}

		_, err := esutil.SearchIndex(ctx, client, IndexName, query)
		if err != nil {
			panic(err)
		}
	},
}

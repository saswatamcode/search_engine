package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	//IndexName is the Elasticsearch index name
	IndexName string
	tWidth    int
)

var rootCmd = &cobra.Command{
	Use:   "searchQuotes",
	Short: "SearchQuotes allows you to index and search quotes from goodreads.com!",
	Long:  "SearchQuotes allows you to index and search quotes from goodreads.com! This is CLI application for demonstrating the utility of Elasticsearch.",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&IndexName, "index", "i", "quotes_index", "Elasticsearch index name")
	tWidth, _, _ = terminal.GetSize(int(os.Stdout.Fd()))
}

// Execute launches the CLI application.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

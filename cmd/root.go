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
	Use:   "quotes",
	Short: "Quotes allows you to index and search quotes from goodreads.com!",
	Long: `Quotes allows you to index and search quotes from goodreads.com!
	index: Indexes quote from goodread.com 
	`,
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

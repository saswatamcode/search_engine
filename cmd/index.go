package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "search",
	Short: "A very fast search_engine",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yo")
	},
}

//Execute is the command executor
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

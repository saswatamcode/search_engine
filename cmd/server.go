package cmd

import (
	"search_engine/server"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts server for searching quotes",
	Run: func(cmd *cobra.Command, args []string) {
		server.StartServer()
	},
}

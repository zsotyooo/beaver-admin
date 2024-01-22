package cmd

import (
	"api/internal/bootstrap"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  "Serve app on dev server",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	bootstrap.Serve()
}

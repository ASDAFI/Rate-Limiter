package cli

import (
	"github.com/spf13/cobra"
	"log"
	"service/src/infrastructure/server"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving rate limiter service.",
	Long:  `Serving rate limiter service.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() {
	go server.RunInsecureGRPCServer()
	err := server.RunInsecureHTTPServer()
	if err != nil {
		log.Fatal(err)
	}
}

package cli

import (
	"github.com/spf13/cobra"
	"log"
	"service/src/infrastructure/db/migration"
)

var migrate = &cobra.Command{
	Use: "migrate",
	Run: func(cobraCmd *cobra.Command, args []string) {
		if err := migration.MigrateDB(); err != nil {
			log.Panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrate)
}

package cli

import (
	"github.com/spf13/cobra"
	"log"
	"service/src/infrastructure/cache"
	"service/src/infrastructure/db"
	"service/src/users"
)

var createObject = &cobra.Command{
	Use:   "create [src]",
	Short: "Run farm for create db objects.",
	Long:  `Run farm for create db objects`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {

		case "user":
			dbProvider := db.PostgresDBProvider
			cacheProvider := cache.RedisCacheProvider

			usersRepo := users.NewUserRepository(dbProvider, cacheProvider)

			userCHM := users.NewCommandHandler(usersRepo)

			command := users.CreateUserCommand{

				Username:  args[1],
				Password:  args[2],
				Email:     args[3],
				FirstName: args[4],
			}

			_, err := userCHM.CreateUser(cmd.Context(), command)
			if err != nil {
				log.Panic(err)
			}

		default:
			log.Panic("ops !")
		}
	},
}

func init() {
	rootCmd.AddCommand(createObject)
}

package cli

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"service/configs"
	"service/src/infrastructure/cache"
	"service/src/infrastructure/db"
	"service/src/users"
)

var rateLimiterCmd = &cobra.Command{
	Use:   "rate_limits",
	Short: "Print the rate limits info",
	Long:  `Print complete rate limits info.`,
	Run: func(cmd *cobra.Command, args []string) {
		logRateLimits()
	},
}

func init() {
	rootCmd.AddCommand(rateLimiterCmd)
}

func setRateLimits() {
	dbProvider := db.PostgresDBProvider
	cacheProvider := cache.RedisCacheProvider

	usersRepo := users.NewUserRepository(dbProvider, cacheProvider)

	userCHM := users.NewCommandHandler(usersRepo)

	for _, rateLimit := range configs.Config.RateLimits {
		command := users.SetRateLimitCommand{RequestsPerMinute: uint(rateLimit.RequestsPerMinute), RPCName: rateLimit.RPCName}
		err := userCHM.SetRateLimit(context.Background(), command)
		if err != nil {
			log.Fatal("Error setting rate limit:", err)
		}

	}
}
func getRateLimitsAsString() string {
	result := "Rate limits: \n"
	for _, rateLimit := range configs.Config.RateLimits {
		result += fmt.Sprintf("Name: %v \t RequestsPerMinute: %v \n", rateLimit.RPCName, rateLimit.RequestsPerMinute)
	}
	return result
}

func logRateLimits() {
	log.Print(getRateLimitsAsString())
}

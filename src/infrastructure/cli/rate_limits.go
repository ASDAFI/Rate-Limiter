package cli

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"service/configs"
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

func getRateLimitsAsString() string {
	result := "Rate limits: \n"
	for _, rateLimit := range configs.Config.RateLimits {
		result += fmt.Sprintf("Name: %v \t RequestsPerSecond: %v \n", rateLimit.RPCName, rateLimit.RequestsPerSecond)
	}
	return result
}

func logRateLimits() {
	log.Print(getRateLimitsAsString())
}

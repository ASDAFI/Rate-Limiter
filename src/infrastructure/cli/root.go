package cli

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "rate limiter",
		Short: "rate limiter service.",
		Long:  `rate limiter service middleware.`,
	}
)

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "service-configs", "config file")
	rootCmd.PersistentFlags().StringP("author", "a", "Ali Sadafi", "bale@gmail.com")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(setupConfig)
	cobra.OnInitialize(setRateLimits)
}

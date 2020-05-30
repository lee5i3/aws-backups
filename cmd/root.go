package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	region  string
	profile string

	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	rootCmd.PersistentFlags().StringVarP(&profile, "profile", "p", "default", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&region, "region", "r", "us-east-1", "name of license for the project")
}

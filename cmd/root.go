package cmd

import (
	"github.com/spf13/cobra"
)

var (
	gitCommit string

	// Used for flags.
	region  string
	profile string

	rootCmd = &cobra.Command{
		Use:   "aws-backups",
		Short: "A generator for Cobra based Applications",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	rootCmd.PersistentFlags().StringVarP(&profile, "profile", "p", "default", "AWS profile to use for connecting")
	rootCmd.PersistentFlags().StringVarP(&region, "region", "r", "us-east-1", "AWS region")
}

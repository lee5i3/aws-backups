package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(cleanupCmd)
}

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup old backups",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Cleaning up")
	},
}

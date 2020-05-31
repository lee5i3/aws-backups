package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(backupCmd)
}

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup all servers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Backing up")
	},
}

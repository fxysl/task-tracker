package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task-tracker",
	Short: "CLI Ngan Saukur Tugas",
}

func Execute() error {
	rootCmd.AddCommand(NewAddCmd())
	rootCmd.AddCommand(NewUpdateCmd())
	rootCmd.AddCommand(NewDeleteCmd())
	rootCmd.AddCommand(NewMarkInProgressCmd())
	rootCmd.AddCommand(NewMarkDoneCmd())
	rootCmd.AddCommand(NewListCmd())
	return rootCmd.Execute()
}

package cmd

import (
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Daptar sakabeh pagawean",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			statusFilter := ""
			if len(args) == 0 {
				statusFilter = args[0]
			}
			task.ListTasks(statusFilter)
			return nil
		},
	}
	return listCmd
}

package cmd

import (
	"errors"
	"fmt"
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add [description]",
		Short: "Nambahkeun pagawean anyar",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := task.AddTask(args[0])
			if err != nil {
				return errors.New("gagal nambahkeun pagawean: " + err.Error())
			}
			fmt.Printf("Pagawean hasil ditambahkeun (ID: %d)\n", id.ID)
			return nil
		},
	}
	return addCmd
}

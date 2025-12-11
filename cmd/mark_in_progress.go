package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

func NewMarkInProgressCmd() *cobra.Command {
	markInProgressCmd := &cobra.Command{
		Use:   "mark-in-progress [id]",
		Short: "Nandaan pagawean salaku keur digawean",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return errors.New("ID kudu angka sujang")
			}
			err = task.MarkInProgress(id)
			if err != nil {
				return err
			}
			fmt.Printf("Pagawean %d hasil ditandaan salaku keur digawean\n", id)
			return nil
		},
	}
	return markInProgressCmd
}

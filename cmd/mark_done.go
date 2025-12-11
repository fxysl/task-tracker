package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

func NewMarkDoneCmd() *cobra.Command {
	markDoneCmd := &cobra.Command{
		Use:   "mark-done [id]",
		Short: "Nandaan pagawean anu geus beres",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return errors.New("ID kudu angka sujang")
			}
			err = task.MarkDone(id)
			if err != nil {
				return err
			}
			fmt.Printf("Pagawean %d hasil ditandaan salaku geus beres\n", id)
			return nil
		},
	}
	return markDoneCmd
}

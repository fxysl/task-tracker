package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

func NewUpdateCmd() *cobra.Command {
	UpdateCmd := &cobra.Command{
		Use:   "update [id] [new description]",
		Short: "Nganyarkeun deui pagawean",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return errors.New("ID kudu angka sujang")
			}
			err = task.UpdateTask(id, args[1])
			if err != nil {
				return err
			}
			fmt.Printf("Pagawean %d nganyarkuen berhasil\n", id)
			return nil
		},
	}
	return UpdateCmd
}

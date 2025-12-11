package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	DeleteCmd := &cobra.Command{
		Use:   "delete [id]",
		Short: "Ngahapus pagawean",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return errors.New("ID kudu angka sujang")
			}
			err = task.DeleteTask(id)
			if err != nil {
				return err
			}
			fmt.Printf("Pagawean %d hasil dihapus\n", id)
			return nil
		},
	}
	return DeleteCmd
}

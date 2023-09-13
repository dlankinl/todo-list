package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "delete task",
	Long:  "Delete task by name from your To-Do list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting task...")
	},
}

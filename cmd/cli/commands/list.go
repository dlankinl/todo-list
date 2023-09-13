package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCommand = &cobra.Command{
	Use:   "list",
	Short: "list your tasks",
	Long:  "List all of your tasks in To-Do list application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Printing list of tasks...")
	},
}

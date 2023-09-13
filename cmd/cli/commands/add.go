package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AddCommand = &cobra.Command{
	Use:   "add",
	Short: "add task",
	Long:  "add task by to your To-Do list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding task...")
	},
}

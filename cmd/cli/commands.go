package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "To-Do list application",
	Long:  "To-Do list is the application that can help you to be more efficient!",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is a first command!")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

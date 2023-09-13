package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RegisterCommand = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Long:  "Register a new user with the To-Do list application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Registering a new user...")
	},
}

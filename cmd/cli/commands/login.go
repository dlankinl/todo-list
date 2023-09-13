package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var LoginCommand = &cobra.Command{
	Use:   "login",
	Short: "login with your password",
	Long:  "Login with your password with the To-Do list application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Logining...")
	},
}

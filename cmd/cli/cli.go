package cli

import "todo/cmd/cli/commands"

func init() {
	rootCmd.AddCommand(commands.RegisterCommand)
	rootCmd.AddCommand(commands.LoginCommand)
	rootCmd.AddCommand(commands.ListCommand)
	rootCmd.AddCommand(commands.DeleteCommand)
}

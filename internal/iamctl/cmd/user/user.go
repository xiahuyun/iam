package user

import (
	"github.com/spf13/cobra"
)

// NewUserCommand create the user command
func NewUserCommand() *cobra.Command {
	var userCmd = &cobra.Command{
		Use:   "user",
		Short: "User management",
		Long:  `Manage the IAM system users, including create, get, update and delete users.`,
	}

	// add user sub commands
	userCmd.AddCommand(NewUserCreateCommand())
	userCmd.AddCommand(NewUserListCommand())
	userCmd.AddCommand(NewUserGetCommand())
	userCmd.AddCommand(NewUserUpdateCommand())
	userCmd.AddCommand(NewUserDeleteCommand())

	return userCmd
}

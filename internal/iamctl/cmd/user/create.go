package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	userName     string
	userPassword string
	userPhone    string
)

// NewUserCreateCommand create the user create command
func NewUserCreateCommand() *cobra.Command {
	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new user",
		Long:  `Create a new user in the IAM system.`,
		Run: func(cmd *cobra.Command, args []string) {
			// validate the create user command
			if userName == "" {
				fmt.Println("Username is required")
				cmd.Help()
				return
			}
			if userPhone == "" {
				fmt.Println("Phone is required")
				cmd.Help()
				return
			}
			if userPassword == "" {
				fmt.Println("Password is required")
				cmd.Help()
				return
			}

			// 实现创建用户的逻辑
			fmt.Printf("创建用户: %s (手机号: %s)\n", userName, userPhone)
			// 这里应该调用相应的服务来创建用户
		},
	}

	// create user command flags
	createCmd.Flags().StringVarP(&userName, "name", "n", "", "用户名")
	createCmd.Flags().StringVarP(&userPhone, "phone", "p", "", "用户手机号")
	createCmd.Flags().StringVarP(&userPassword, "password", "w", "", "用户密码")

	// set required flags
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("phone")
	createCmd.MarkFlagRequired("password")

	return createCmd
}

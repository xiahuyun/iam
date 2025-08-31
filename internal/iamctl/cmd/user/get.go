package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewUserGetCommand 创建获取用户命令
func NewUserGetCommand() *cobra.Command {
	var getCmd = &cobra.Command{
		Use:   "get [用户名]",
		Short: "Get user details",
		Long:  `Get details of the specified user.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			userName := args[0]
			// 实现获取用户详情的逻辑
			fmt.Printf("Get user '%s' details:\n", userName)
			fmt.Println("--------------------")
			fmt.Printf("Username: %s\n", userName)
			fmt.Printf("Email: example@example.com\n")
			fmt.Printf("Role: Admin\n")
			fmt.Printf("Create Time: 2023-01-01 10:00:00\n")
			fmt.Printf("Last Login: 2023-06-01 15:30:00\n")
			// 这里应该调用相应的服务来获取用户详情
		},
	}

	return getCmd
}

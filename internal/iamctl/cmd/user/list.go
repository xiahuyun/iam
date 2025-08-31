package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewUserListCommand 创建列出用户命令
func NewUserListCommand() *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all users",
		Long:  `List all users in the IAM system.`,
		Run: func(cmd *cobra.Command, args []string) {
			// 实现列出用户的逻辑
			fmt.Println("List all users:")
			fmt.Println("--------------------")
			fmt.Println("ID | Username | Email | Role")
			fmt.Println("--------------------")
			// 这里应该调用相应的服务来获取用户列表
		},
	}

	return listCmd
}

package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewUserListCommand 创建列出用户命令
func NewUserListCommand() *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "列出所有用户",
		Long:  `列出 IAM 系统中的所有用户。`,
		Run: func(cmd *cobra.Command, args []string) {
			// 实现列出用户的逻辑
			fmt.Println("列出所有用户:")
			fmt.Println("--------------------")
			fmt.Println("ID | 用户名 | 邮箱 | 角色")
			fmt.Println("--------------------")
			// 这里应该调用相应的服务来获取用户列表
		},
	}

	return listCmd
}

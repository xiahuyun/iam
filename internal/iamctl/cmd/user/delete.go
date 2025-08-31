package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewUserDeleteCommand 创建删除用户命令
func NewUserDeleteCommand() *cobra.Command {
	var deleteCmd = &cobra.Command{
		Use:   "delete [用户名]",
		Short: "删除用户",
		Long:  `从 IAM 系统中删除指定的用户。`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			userName := args[0]
			// 实现删除用户的逻辑
			fmt.Printf("删除用户 '%s'\n", userName)
			// 这里应该调用相应的服务来删除用户
		},
	}

	return deleteCmd
}

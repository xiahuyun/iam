package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewUserGetCommand 创建获取用户命令
func NewUserGetCommand() *cobra.Command {
	var getCmd = &cobra.Command{
		Use:   "get [用户名]",
		Short: "获取用户详情",
		Long:  `获取指定用户名的用户详情。`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			userName := args[0]
			// 实现获取用户详情的逻辑
			fmt.Printf("获取用户 '%s' 详情:\n", userName)
			fmt.Println("--------------------")
			fmt.Printf("用户名: %s\n", userName)
			fmt.Printf("邮箱: example@example.com\n")
			fmt.Printf("角色: Admin\n")
			fmt.Printf("创建时间: 2023-01-01 10:00:00\n")
			fmt.Printf("最后登录: 2023-06-01 15:30:00\n")
			// 这里应该调用相应的服务来获取用户详情
		},
	}

	return getCmd
}

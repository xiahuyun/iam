package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewUserUpdateCommand 创建更新用户命令
func NewUserUpdateCommand() *cobra.Command {
	var updateCmd = &cobra.Command{
		Use:   "update [用户名]",
		Short: "更新用户信息",
		Long:  `更新指定用户名的用户信息。`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			userName := args[0]
			// 实现更新用户信息的逻辑
			fmt.Printf("更新用户 '%s' 信息\n", userName)
			// 检查哪些标志被设置了
			if userPhone != "" {
				fmt.Printf("新手机号: %s\n", userPhone)
			}
			if userPassword != "" {
				fmt.Println("密码已更新")
			}

			// 这里应该调用相应的服务来更新用户信息
		},
	}

	// 添加更新用户命令的标志
	updateCmd.Flags().StringVarP(&userPhone, "phone", "p", "", "用户手机号")
	updateCmd.Flags().StringVarP(&userPassword, "password", "w", "", "用户密码")

	return updateCmd
}

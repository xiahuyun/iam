package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

type UpdateOption struct {
	Username string
	Phone    string
	Password string
}

func NewUpdateOption() *UpdateOption {
	return &UpdateOption{}
}

func (o *UpdateOption) Validate() error {
	if o.Username == "" {
		return fmt.Errorf("username is required")
	}
	return nil
}

// NewUserUpdateCommand 创建更新用户命令
func NewUserUpdateCommand() *cobra.Command {
	o := NewUpdateOption()
	var updateCmd = &cobra.Command{
		Use:   "update [用户名]",
		Short: "Update user information",
		Long:  `Update the information of the specified user.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Validate(); err != nil {
				return err
			}

			return nil
		},
	}

	// update user command flags
	updateCmd.Flags().StringVarP(&o.Phone, "phone", "p", "", "User phone number")
	updateCmd.Flags().StringVarP(&o.Password, "password", "w", "", "User password")

	return updateCmd
}

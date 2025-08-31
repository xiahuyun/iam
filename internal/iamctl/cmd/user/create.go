package user

import (
	"context"
	"fmt"

	v1 "github.com/marmotedu/api/apiserver/v1"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	apiclientv1 "github.com/marmotedu/marmotedu-sdk-go/marmotedu/service/iam/apiserver/v1"
	"github.com/marmotedu/marmotedu-sdk-go/tools/clientcmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CreateOption struct {
	User   *v1.User
	Client apiclientv1.APIV1Interface
}

func NewCreateOption() *CreateOption {
	return &CreateOption{
		User: &v1.User{},
	}
}

// Validate validates the create option.
func (o *CreateOption) Validate() error {
	if errs := o.User.Validate(); len(errs) > 0 {
		return errs.ToAggregate()
	}

	return nil
}

// Complete completes all the required options.
func (o *CreateOption) Complete() error {
	var err error

	config := clientcmd.NewConfig()
	if err = viper.Unmarshal(&config); err != nil {
		return err
	}

	clientConfig, err := clientcmd.NewClientConfigFromConfig(config).ClientConfig()
	if err != nil {
		return err
	}

	o.Client, err = apiclientv1.NewForConfig(clientConfig)
	if err != nil {
		return err
	}

	return nil
}

func (o *CreateOption) Run() error {
	user, err := o.Client.Users().Create(context.Background(), o.User, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	fmt.Printf("User %s created successfully\n", user.Name)

	return nil
}

// NewUserCreateCommand create the user create command
func NewUserCreateCommand() *cobra.Command {
	o := NewCreateOption()

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new user",
		Long:  `Create a new user in the IAM system.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}

			if err := o.Validate(); err != nil {
				return err
			}

			return o.Run()
		},
	}

	// create user command flags
	createCmd.Flags().StringVarP(&o.User.Name, "username", "u", "", "user name")
	createCmd.Flags().StringVar(&o.User.Phone, "phone", "", "user phone number")
	createCmd.Flags().StringVarP(&o.User.Password, "password", "p", "", "user password")
	createCmd.Flags().StringVarP(&o.User.Nickname, "nickname", "n", "", "user nickname")
	createCmd.Flags().StringVarP(&o.User.Email, "email", "e", "", "user email")

	return createCmd
}

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xiahuyun/iam/internal/iamctl/cmd/user"
)

// NewRootCommand create the root command
func NewRootCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "iamctl",
		Short: "IAM command line tool",
		Long: `iamctl is a command line tool for managing IAM system.

It provides user management, secret management and policy management for authentication and authorization.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
			} else {
				fmt.Println("Unknown command:", args[0])
				cmd.Help()
				os.Exit(1)
			}
		},
	}

	flags := rootCmd.PersistentFlags()

	// add global flags
	flags.StringP("config", "c", "", "Config file path")
	flags.BoolP("verbose", "v", false, "Enable verbose output")

	// add sub commands
	rootCmd.AddCommand(NewUserCommand())

	_ = viper.BindPFlags(flags)
	cobra.OnInitialize(initConfig)

	return rootCmd
}

// initConfig initialize the iamctl config file
func initConfig() {
	configPath := viper.GetString("config")
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("$HOME/.iamctl")
		viper.AddConfigPath(".")
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("iamctl")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}
}

// NewUserCommand create the user command
func NewUserCommand() *cobra.Command {
	var userCmd = &cobra.Command{
		Use:   "user",
		Short: "User management",
		Long:  `Manage IAM system users, including create, query, update and delete users.`,
	}

	// add user sub commands
	userCmd.AddCommand(user.NewUserCreateCommand())
	userCmd.AddCommand(user.NewUserListCommand())
	userCmd.AddCommand(user.NewUserGetCommand())
	userCmd.AddCommand(user.NewUserUpdateCommand())
	userCmd.AddCommand(user.NewUserDeleteCommand())

	return userCmd
}

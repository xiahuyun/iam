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
		Run: runHelp,
	}

	flags := rootCmd.PersistentFlags()

	// add global flags
	flags.StringP("config", "c", "", "Config file path")
	flags.BoolP("verbose", "v", false, "Enable verbose output")

	// add sub commands
	rootCmd.AddCommand(user.NewUserCommand())

	_ = viper.BindPFlags(flags)
	cobra.OnInitialize(initConfig)

	return rootCmd
}

func runHelp(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}

// initConfig initialize the iamctl config file
func initConfig() {
	configPath := viper.GetString("config")
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.AddConfigPath("$HOME/.iamctl")
		viper.AddConfigPath(".")

		// @TODO: remove the default config path as soon as possible
		// the config path exists only for test purpose
		viper.AddConfigPath("./configs")

		viper.SetConfigName("iamctl.yaml")
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

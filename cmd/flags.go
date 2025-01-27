package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type preRunFunc func(cmd *cobra.Command, args []string)

func flagsAppID_Token(cmd *cobra.Command) preRunFunc {
	cmd.Flags().String("app-id", "", "GitHub App ID")
	viper.BindPFlag("app-id", cmd.Flags().Lookup("app-id"))
	viper.BindEnv("app-id", "GITHUB_APP_ID")

	cmd.Flags().String("private-key", "", "Path to the private key file or the private key itself")
	viper.BindPFlag("private-key", cmd.Flags().Lookup("private-key"))
	viper.BindEnv("private-key", "GITHUB_APP_PRIVATE_KEY")

	cmd.MarkFlagRequired("app-id")
	cmd.MarkFlagRequired("private-key")

	preRunFn := func(cmd *cobra.Command, args []string) {
		if viper.IsSet("app-id") {
			cmd.Flags().Set("app-id", viper.GetString("app-id"))
		}
		if viper.IsSet("private-key") {
			cmd.Flags().Set("private-key", viper.GetString("private-key"))
		}
	}
	return preRunFn
}

func requiredFlag(cmd *cobra.Command, name string, value string, usage string, envVar string) preRunFunc {
	cmd.Flags().String(name, value, usage)
	viper.BindPFlag(name, cmd.Flags().Lookup(name))
	viper.BindEnv(name, envVar)

	cmd.MarkFlagRequired(name)

	preRunFn := func(cmd *cobra.Command, args []string) {
		if viper.IsSet(name) {
			cmd.Flags().Set(name, viper.GetString(name))
		}
	}
	return preRunFn
}

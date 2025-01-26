package cmd

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newRootCmd() *cobra.Command {
	var rootCmd = cobra.Command{
		Use:   "gh-app-token",
		Short: "A cli utility to manage GitHub App tokens",
		Long:  `A cli utility to create and revoke Github App installation tokens`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
			}
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if viper.IsSet("app-id") {
				_ = cmd.Flags().Set("app-id", viper.GetString("app-id"))
			}
			if viper.IsSet("private-key") {
				_ = cmd.Flags().Set("private-key", viper.GetString("private-key"))
			}
			return nil
		},
	}

	viper.AutomaticEnv()

	rootCmd.PersistentFlags().String("app-id", "", "GitHub App ID")
	viper.BindPFlag("app-id", rootCmd.PersistentFlags().Lookup("app-id"))
	viper.BindEnv("app-id", "GITHUB_APP_ID")

	rootCmd.PersistentFlags().String("private-key", "", "Path to the private key file or the private key itself")
	viper.BindPFlag("private-key", rootCmd.PersistentFlags().Lookup("private-key"))
	viper.BindEnv("private-key", "GITHUB_APP_PRIVATE_KEY")

	rootCmd.MarkPersistentFlagRequired("app-id")
	rootCmd.MarkPersistentFlagRequired("private-key")

	return &rootCmd
}

func Execute(args []string, out io.Writer, err io.Writer) error {
	rootCmd := newRootCmd()
	rootCmd.AddCommand(newCreateCmd())
	rootCmd.AddCommand(newRevokeCmd())
	rootCmd.AddCommand(newInstallationCmd())

	rootCmd.SetArgs(args)
	rootCmd.SetOut(out)
	rootCmd.SetErr(err)

	return rootCmd.Execute()
}

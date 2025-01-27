package cmd

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/n26/gh-app-token/internal/github"
)

type GithubClientFactory func(string) github.GithubClient

func newRootCmd(githubClientFactory GithubClientFactory) *cobra.Command {
	var rootCmd = cobra.Command{
		Use:   "gh-app-token",
		Short: "A cli utility to manage GitHub App tokens",
		Long:  `A cli utility to create and revoke Github App installation tokens`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
			}
		},
	}

	rootCmd.AddCommand(newCreateCmd(githubClientFactory))
	rootCmd.AddCommand(newRevokeCmd(githubClientFactory))
	rootCmd.AddCommand(newInstallationCmd())

	return &rootCmd
}

func Execute(args []string, out io.Writer, err io.Writer) error {
	viper.AutomaticEnv()

	rootCmd := newRootCmd(github.NewClient)

	rootCmd.SetArgs(args)
	rootCmd.SetOut(out)
	rootCmd.SetErr(err)

	return rootCmd.Execute()
}

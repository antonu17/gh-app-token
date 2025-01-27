package cmd

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/n26/gh-app-token/internal/github"
	"github.com/n26/gh-app-token/internal/jwt"
)

type GithubClientFactory func(token string) github.GithubClient
type JWTTokenFactory func(issuer string, privateKey string) (string, error)

func newRootCmd(githubClientFactory GithubClientFactory, jwtTokenFactory JWTTokenFactory) *cobra.Command {
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

	rootCmd.AddCommand(newCreateCmd(githubClientFactory, jwtTokenFactory))
	rootCmd.AddCommand(newRevokeCmd(githubClientFactory))
	rootCmd.AddCommand(newInstallationsCmd(githubClientFactory, jwtTokenFactory))

	return &rootCmd
}

func Execute(args []string, out io.Writer, err io.Writer) error {
	viper.AutomaticEnv()

	rootCmd := newRootCmd(github.NewClient, jwt.NewToken)

	rootCmd.SetArgs(args)
	rootCmd.SetOut(out)
	rootCmd.SetErr(err)

	return rootCmd.Execute()
}

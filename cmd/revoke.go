package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

func newRevokeCmd(githubClientFactory GithubClientFactory) *cobra.Command {
	const defaultTimeout = 10 * time.Second
	var revokeCmd = cobra.Command{
		Use:   "revoke",
		Short: "Revoke an existing GitHub App token",
		Run: func(cmd *cobra.Command, args []string) {
			token, _ := cmd.Flags().GetString("token")

			gh := githubClientFactory(token)
			err := gh.RevokeInstallationToken()
			if err != nil {
				cmd.PrintErrln("error revoking token", err)
				return
			}

			cmd.Println("token sucessfully revoked")
		},
	}

	tokenPreRunCheck := requiredFlag(&revokeCmd, "token", "", "Installation token to revoke", "GITHUB_TOKEN")
	revokeCmd.PreRun = func(cmd *cobra.Command, args []string) {
		tokenPreRunCheck(cmd, args)
	}
	return &revokeCmd
}

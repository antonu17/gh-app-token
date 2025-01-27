package cmd

import (
	"os"
	"time"

	"github.com/n26/gh-app-token/internal/jwt"
	"github.com/spf13/cobra"
)

func newCreateCmd(githubClientFactory GithubClientFactory) *cobra.Command {
	const defaultTimeout = 10 * time.Second
	var createCmd = cobra.Command{
		Use:   "create",
		Short: "Create a new GitHub App token",
		Run: func(cmd *cobra.Command, args []string) {
			appID, _ := cmd.Flags().GetString("app-id")
			privateKey, _ := cmd.Flags().GetString("private-key")

			if _, err := os.Stat(privateKey); err == nil {
				content, err := os.ReadFile(privateKey)
				if err != nil {
					cmd.PrintErrln("error reading private key file:", err)
					return
				}
				privateKey = string(content)
			}

			appToken, err := jwt.NewToken(appID, privateKey)
			if err != nil {
				cmd.PrintErrln("error generating JWT token for Github App:", err)
				return
			}

			gh := githubClientFactory(appToken)

			installationID, err := gh.GetInstallationID()
			if err != nil {
				cmd.PrintErrln("error getting installation id:", err)
				return
			}

			installationToken, err := gh.CreateInstallationToken(installationID)
			if err != nil {
				cmd.PrintErrln("error generating installation token for Github App:", err)
				return
			}

			cmd.Println(installationToken)
		},
	}

	flagsPreRunCheck := flagsAppID_Token(&createCmd)

	createCmd.PreRun = func(cmd *cobra.Command, args []string) {
		flagsPreRunCheck(cmd, args)
	}

	return &createCmd
}

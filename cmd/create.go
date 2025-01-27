package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func newCreateCmd(githubClientFactory GithubClientFactory, jwtTokenFactory JWTTokenFactory) *cobra.Command {
	const defaultTimeout = 10 * time.Second
	var createCmd = cobra.Command{
		Use:   "create",
		Short: "Create a new GitHub App token",
		RunE: func(cmd *cobra.Command, args []string) error {
			appID, _ := cmd.Flags().GetString("app-id")
			privateKey := loadPrivateKey(cmd)

			appToken, err := jwtTokenFactory(appID, privateKey)
			if err != nil {
				err := fmt.Errorf("error generating JWT token for Github App: %w", err)
				return err
			}

			gh := githubClientFactory(appToken)

			installationID, err := gh.GetInstallationID()
			if err != nil {
				err := fmt.Errorf("error getting installation id: %w", err)
				return err
			}

			installationToken, err := gh.CreateInstallationToken(installationID)
			if err != nil {
				err := fmt.Errorf("error generating installation token for Github App: %w", err)
				return err
			}

			cmd.Println(installationToken)
			return nil
		},
	}

	flagsPreRunCheck := flagsAppID_Token(&createCmd)

	createCmd.PreRun = func(cmd *cobra.Command, args []string) {
		flagsPreRunCheck(cmd, args)
	}
	return &createCmd
}

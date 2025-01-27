package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newInstallationsCmd(githubClientFactory GithubClientFactory, jwtTokenFactory JWTTokenFactory) *cobra.Command {
	var installationsCmd = cobra.Command{
		Use:   "installations",
		Short: "Get GitHub App Installations",
		RunE: func(cmd *cobra.Command, args []string) error {
			appID, _ := cmd.Flags().GetString("app-id")
			privateKey := loadPrivateKey(cmd)

			appToken, err := jwtTokenFactory(appID, privateKey)
			if err != nil {
				err := fmt.Errorf("error generating JWT token for Github App: %w", err)
				return err
			}

			gh := githubClientFactory(appToken)

			installations, err := gh.ListInstallations()
			if err != nil {
				err := fmt.Errorf("error generating JWT token for Github App: %w", err)
				return err
			}
			cmd.Println(installations)
			return nil
		},
	}

	flagsPreRunCheck := flagsAppID_Token(&installationsCmd)

	installationsCmd.PreRun = func(cmd *cobra.Command, args []string) {
		flagsPreRunCheck(cmd, args)
	}
	return &installationsCmd
}

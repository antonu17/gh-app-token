package cmd

import (
	"fmt"
	"os"

	"github.com/n26/gh-app-token/internal/github"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	var createCmd = cobra.Command{
		Use:   "create",
		Short: "Create a new GitHub App token",
		Run: func(cmd *cobra.Command, args []string) {
			appID, _ := cmd.Flags().GetString("app-id")
			privateKey, _ := cmd.Flags().GetString("private-key")

			if _, err := os.Stat(privateKey); err == nil {
				content, err := os.ReadFile(privateKey)
				if err != nil {
					fmt.Println("error reading private key file:", err)
					return
				}
				privateKey = string(content)
			}

			appToken, err := github.NewAppToken(appID, privateKey)
			if err != nil {
				fmt.Println("error generating JWT token for Github App:", err)
				return
			}

			installationID, err := github.GetInstallationID(appToken)
			if err != nil {
				fmt.Println("error getting installation id:", err)
				return
			}

			installationToken, err := github.NewInstallationToken(appToken, installationID)
			if err != nil {
				fmt.Println("error generating installation token for Github App:", err)
				return
			}

			fmt.Println(installationToken)
		},
	}
	return &createCmd
}

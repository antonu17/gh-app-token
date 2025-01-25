package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var revokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: "Revoke an existing GitHub App token",
	Run: func(cmd *cobra.Command, args []string) {
		token := getEnv("TOKEN_ENV", "")
		fmt.Println("Revoked token:", token)
	},
}

func init() {
	rootCmd.AddCommand(revokeCmd)

	revokeCmd.Flags().StringP("token", "t", "", "A flag for the revoke subcommand")
}

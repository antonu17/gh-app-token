package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new GitHub App token",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("App ID:", appID)
		fmt.Println("Private Key:", privateKey)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("token", "t", "", "A flag for the create subcommand")
}

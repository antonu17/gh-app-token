package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var installationsCmd = &cobra.Command{
	Use:   "installations",
	Short: "Get GitHub App Installations",
	Run: func(cmd *cobra.Command, args []string) {
		flagValue := getEnv("FLAG_ENV", "")
		fmt.Println("Flag value:", flagValue)
	},
}

func init() {
	rootCmd.AddCommand(installationsCmd)

	installationsCmd.Flags().StringP("flag", "f", "", "A flag for the installations subcommand")
}

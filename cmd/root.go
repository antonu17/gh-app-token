package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (
	appID      string
	privateKey string
)

var rootCmd = &cobra.Command{
	Use:   "gh-app-token",
	Short: "A CLI tool to manage GitHub App tokens",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat(privateKey); err == nil {
			content, err := os.ReadFile(privateKey)
			if err != nil {
				return err
			}
			privateKey = string(content)
		} else if !os.IsNotExist(err) {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func Execute(args []string, out io.Writer) error {
	rootCmd.SetArgs(args)
	rootCmd.SetOut(out)
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&appID, "app-id", "", "GitHub App ID")
	rootCmd.PersistentFlags().StringVar(&privateKey, "private-key", "", "Path to the private key file or the private key itself")
	rootCmd.MarkPersistentFlagRequired("app-id")
	rootCmd.MarkPersistentFlagRequired("private-key")
}

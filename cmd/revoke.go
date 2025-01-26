package cmd

import (
	"github.com/spf13/cobra"
)

func newRevokeCmd() *cobra.Command {
	var revokeCmd = cobra.Command{
		Use:   "revoke",
		Short: "Revoke an existing GitHub App token",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return &revokeCmd
}

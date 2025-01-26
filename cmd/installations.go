package cmd

import (
	"github.com/spf13/cobra"
)

func newInstallationCmd() *cobra.Command {
	var installationsCmd = cobra.Command{
		Use:   "installations",
		Short: "Get GitHub App Installations",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return &installationsCmd
}

package cmd

import (
	"fmt"
	"github.com/sibprogrammer/uws/internal/server"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "uws",
	Short: "Universal Web Server for development purposes",
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.Run()
	},
}

func Execute() {
	rootCmd.AddCommand(
		versionCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

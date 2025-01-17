package cmd

import (
	"github.com/spf13/cobra"
)

var (

	rootCmd = &cobra.Command{
		Use:   "IP-Tracker",
		Short: "IP-Tracker CLI ",
		Long: `IP-Tracker is used to track IP address location.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

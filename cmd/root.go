package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "IP-tracker CLI app",
		Long:  "IP-tracker CLI app",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

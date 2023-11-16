package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd prints the version
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "getting the version",
	Long:  "getting the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version: 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

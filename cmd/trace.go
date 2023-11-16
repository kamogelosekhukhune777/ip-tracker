package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  "Trace the IP",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("testing trace sub command")
	},
}

//07:00

func init() {
	rootCmd.AddCommand(traceCmd)
}

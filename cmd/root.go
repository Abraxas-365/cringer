package cmd

import (
	"os"

	"github.com/Abraxas-365/cringer/cmd/tmsg"
	"github.com/spf13/cobra"
)

var from string
var command string
var to []string

var rootCmd = &cobra.Command{
	Use:   "cringer",
	Short: "notify when command finish execution",
	Long:  `notify when command finish execution.`,
	// Run:   func(cmd *cobra.Command, args []string) {}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommandPalletes() {
	rootCmd.AddCommand(tmsg.TmsgCmd)
}
func init() {
	addSubCommandPalletes()
}

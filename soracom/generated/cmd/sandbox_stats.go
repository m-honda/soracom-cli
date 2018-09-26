package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	SandboxCmd.AddCommand(SandboxStatsCmd)
}

// SandboxStatsCmd defines 'stats' subcommand
var SandboxStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: TRCLI("cli.sandbox.stats.summary"),
	Long:  TRCLI(`cli.sandbox.stats.description`),
}

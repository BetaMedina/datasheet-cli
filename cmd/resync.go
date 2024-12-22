package cmd

import (
	"cli/internal/commands"

	"github.com/spf13/cobra"
)

var resyncCmd = &cobra.Command{
	Use:   "datasheet-resync",
	Short: "Run datasheet resync job",
	Long:  `A simple way to running datasheet resync job`,
	Run: func(cmd *cobra.Command, args []string) {
		job, _ := cmd.Flags().GetString("j")
		kind, _ := cmd.Flags().GetString("k")
		commands.RunResync(job, kind, "sync")
	},
}

func init() {
	RootCmd.AddCommand(resyncCmd)
}

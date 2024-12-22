package cmd

import (
	"cli/internal/commands"

	"github.com/spf13/cobra"
)

var renewCmd = &cobra.Command{
	Use:   "datasheet-renew",
	Short: "Run datasheet renew job",
	Long:  `A simple way to running datasheet renew job`,
	Run: func(cmd *cobra.Command, args []string) {
		job, _ := cmd.Flags().GetString("j")
		kind, _ := cmd.Flags().GetString("k")
		commands.RunReprocess(job, kind, "renew")
	},
}

func init() {
	RootCmd.AddCommand(renewCmd)
}

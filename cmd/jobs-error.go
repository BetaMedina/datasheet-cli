package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var jobsErrorCmd = &cobra.Command{
	Use:   "datasheet-jobs-error",
	Short: "Detect a number of jobs with error on database",
	Long:  `A simole way to detect a number of jobs with error on database`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[WIP] - Detecting jobs with error on database")
	},
}

func init() {
	RootCmd.AddCommand(jobsErrorCmd)
}

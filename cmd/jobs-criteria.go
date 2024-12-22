package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var jobsCriteriaCmd = &cobra.Command{
	Use:   "datasheet-jobs-criteria",
	Short: "Get jobs criteria",
	Long:  `This command has the purpose to get jobs criteria`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[WIP] - Detecting jobs criteria")
	},
}

func init() {
	RootCmd.AddCommand(jobsCriteriaCmd)
}

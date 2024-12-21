package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "datasheet-example",
	Short: "An example of datasheet CLI application",
	Long:  `A simple way to interact with datasheet API`,
}

func init() {
	cobra.OnInitialize()
	RootCmd.PersistentFlags().String("j", "", "A jobId parameter")
	RootCmd.PersistentFlags().String("k", "", "A kind parameter")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

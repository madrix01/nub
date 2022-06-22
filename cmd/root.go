package cmd

import (
	"fmt"
	"nub/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nub",
	Short: "Manage your markdown notes using github gist",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to nub")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.MakeError("Unsuccessful running root command")
	}
}

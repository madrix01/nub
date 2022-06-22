package cmd

import (
	"fmt"
	"nub/utils"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create new note",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating new note")
		if len(args) < 1 {
			fmt.Println("[Error] Enter file name")
			os.Exit(1)
		}
		utils.CreateNewNote(args[0])
	},
}

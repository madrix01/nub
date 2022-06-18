package cmd

import (
	"fmt"
	"nub/utils"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit your notes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Editing in progress")
		// OpenEditor()
		utils.UpdateGist()
	},
}

func OpenEditor() {

	cmd := exec.Command("vim")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Finished")
}

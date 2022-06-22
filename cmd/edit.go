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
		fmt.Println("Editing")
		utils.GetFileList()
		edit()
	},
}

func edit() {
	gistNameList := utils.GetFileList()
	gistName := utils.VarPromptSelect("Select gist", gistNameList)
	filePath := utils.CreateTempGist(gistName)
	OpenEditor(filePath)
	utils.UpdateGist(gistName)
}

func OpenEditor(path string) {

	cmd := exec.Command(utils.Cnfg.Editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		utils.MakeError(err.Error())
	}

	fmt.Println("Gist updated " + "https://gist.github.com/" + utils.Cnfg.Username + "/" + utils.Cnfg.GistId)
}

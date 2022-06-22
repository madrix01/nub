package utils

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func VarPromptSelect(label string, items []string) string {

	prompt := promptui.Select{
		Label:        label,
		Items:        items,
		HideHelp:     true,
		HideSelected: true,
	}

	_, result, err := prompt.Run()

	if err != nil {
		MakeError("Error encountered in Prompt.")
	}

	fmt.Println("?", "Template : ", result)
	return result
}

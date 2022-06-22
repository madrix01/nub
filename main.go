package main

import (
	"nub/cmd"
	"nub/utils"
)

func main() {
	cmd.Execute()
}

func init() {
	utils.InitConfig()
	utils.GithubLogin()
}

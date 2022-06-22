package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Config struct {
	Editor     string
	GistId     string
	TempFolder string
	Token      string
	Username   string
}

var Cnfg Config

var configFolderPath string = path.Join(os.Getenv("HOME"), ".config", "nub")
var configFilePath string = path.Join(configFolderPath, "config.json")

func InitConfig() {
	CreateIfNotExists(configFolderPath, 0755)

	if !Exists(configFilePath) {
		var newCnfg Config
		newCnfg.Editor = "vim"
		newCnfg.TempFolder = os.Getenv("TMPDIR")

		b, err := json.MarshalIndent(newCnfg, "", "\t")
		if err != nil {
			MakeError(err.Error())
		}

		CreateFile(configFilePath, string(b))
		return
	}

	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		MakeError(err.Error())
	}

	err = json.Unmarshal(data, &Cnfg)
	if err != nil {
		MakeError(err.Error())
	}
	ConfigValidate(Cnfg)
}

func ConfigValidate(cnfg Config) {
	if strings.Trim(cnfg.GistId, " ") == "" {
		MakeError("Enter Gist ID in " + configFilePath)
	}

	if strings.Trim(cnfg.TempFolder, " ") == "" {
		MakeError("Set a temp path to download the gist")
	}

	if strings.Trim(cnfg.Token, " ") == "" {
		MakeError("Innvalid or empty token")
	}
}

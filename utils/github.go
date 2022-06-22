package utils

import (
	"context"
	"errors"
	"os"
	"path"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

// 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

var GhClient github.Client

func GithubLogin() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: Cnfg.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	GhClient = *client
}

func getGistContent(id, fileName string) (string, error) {

	ctx := context.Background()
	gist, _, err := GhClient.Gists.Get(ctx, id)

	if err != nil {
		return "", errors.New("doesn't have string")
	}

	gistContent := *gist.Files[github.GistFilename(fileName)].Content
	return gistContent, nil
}

// Gist file naming goes like -> g_id.md
func CreateTempGist(gistName string) string {
	tmpPath := os.Getenv("TMPDIR")
	if tmpPath == "" {
		tmpPath = "/tmp"
	}

	fileName := "gist_" + Cnfg.GistId + ".md"
	filePath := path.Join(tmpPath, fileName)

	content, err := getGistContent(Cnfg.GistId, gistName)

	if err != nil {
		MakeError(err.Error())
	}

	CreateFile(filePath, content)
	return filePath
}

func UpdateGist(gistName string) {
	ctx := context.Background()

	tmpPath := os.Getenv("TMPDIR")
	if tmpPath == "" {
		tmpPath = "/tmp"
	}

	fileName := "gist_" + Cnfg.GistId + ".md"
	filePath := path.Join(tmpPath, fileName)

	updatedContent, err := ReadFile(filePath)

	if err != nil {
		MakeError(err.Error())
	}

	gist := new(github.Gist)
	gist.Files = make(map[github.GistFilename]github.GistFile)
	gist.Files[github.GistFilename(gistName)] = github.GistFile{Content: &updatedContent}

	GhClient.Gists.Edit(ctx, Cnfg.GistId, gist)
}

func GetFileList() []string {
	ctx := context.Background()

	gist, _, err := GhClient.Gists.Get(ctx, Cnfg.GistId)

	if err != nil {
		MakeError(err.Error())
	}

	gistMap := gist.Files
	keys := []string{}
	for k := range gistMap {
		keys = append(keys, string(k))
	}

	return keys
}

func CreateNewNote(gistName string) {
	ctx := context.Background()

	emptyString := gistName
	gist := new(github.Gist)
	gist.Files = make(map[github.GistFilename]github.GistFile)
	gist.Files[github.GistFilename(gistName)] = github.GistFile{Content: &emptyString}

	GhClient.Gists.Edit(ctx, Cnfg.GistId, gist)
}

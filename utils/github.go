package utils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

var id string = "f37c4ae65655ffb200ff7f4c82bf0b28"

// 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
// func GithubLogin() {
// 	fmt.Println(os.Getenv("GITHUB_TOKEN"))
// 	ctx := context.Background()
// 	ts := oauth2.StaticTokenSource(
// 		&oauth2.Token{AccessToken: "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"},
// 	)
// 	tc := oauth2.NewClient(ctx, ts)

// 	client := github.NewClient(tc)

// 	// list all repositories for the authenticated user
// 	repos, _, err := client.Gists.Get(ctx, id)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(*repos.Files["test.md"].Content)
// }

var GhClient github.Client

func GithubLogin() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	GhClient = *client
}

func GetGistContent(id string) (string, error) {
	// ctx := context.Background()
	// ts := oauth2.StaticTokenSource(
	// 	&oauth2.Token{AccessToken: "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"},
	// )
	// tc := oauth2.NewClient(ctx, ts)

	// client := github.NewClient(tc)

	ctx := context.Background()
	gist, _, err := GhClient.Gists.Get(ctx, id)

	if err != nil {
		return "", errors.New("doesn't have string")
	}

	gistContent := *gist.Files["test.md"].Content
	return gistContent, nil
}

// Gist file naming goes like -> g_id.md
func CreateTempGist() {
	tmpPath := os.Getenv("TMPDIR")
	if tmpPath == "" {
		tmpPath = "/tmp"
	}

	fileName := "gist_" + id
	filePath := path.Join(tmpPath, fileName)

	content, err := GetGistContent(id)

	if err != nil {
		log.Fatal(err)
	}

	CreateFile(filePath, content)
}

func UpdateGist() {
	ctx := context.Background()

	tmpPath := os.Getenv("TMPDIR")
	if tmpPath == "" {
		tmpPath = "/tmp"
	}

	fileName := "gist_" + id
	filePath := path.Join(tmpPath, fileName)

	updatedContent, err := ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updatedContent)

	gist := new(github.Gist)
	gist.Files = make(map[github.GistFilename]github.GistFile)
	gist.Files["test.md"] = github.GistFile{Content: &updatedContent}
	// gist := &github.Gist{
	// 	Files: make(map[github.GistFilename]github.GistFile),
	// }

	fmt.Println(gist)
	GhClient.Gists.Edit(ctx, id, gist)
}

package gitignore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type file struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Type string `json:"type"`
	Sha  string `json:"sha"`
	Url  string `json:"url"`
}

type tree struct {
	Sha  string `json:"sha"`
	Url  string `json:"url"`
	Tree []file `json:"tree"`
}

var github = map[string]string{
	"rawUrl":   "https://raw.githubusercontent.com",
	"reposApi": "https://api.github.com/repos",
	"repo":     "github/gitignore",
	"branch":   "master",
}

func Get(query string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/%s/%s.gitignore", github["rawUrl"], github["repo"], github["branch"], query)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// TODO prepend header to bytes
	var ignore []byte
	if resp.StatusCode == http.StatusOK {
		ignore, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		return ignore, err
	}

	return nil, fmt.Errorf("could not find gitignore for %s", query)
}

func List() []string {
	return getAllGitignores()
}

func Search(query string) ([]string, error) {
	gitignores := getAllGitignores()

	var ignores []string
	for _, file := range gitignores {
		if strings.Contains(file, query) {
			ignores = append(ignores, file)
		}

	}
	if len(ignores) == 0 {
		return nil, fmt.Errorf("no gitignore found for: %s", query)
	}
	return ignores, nil
}

func getAllGitignores() []string {
	url := fmt.Sprintf("%s/%s/git/trees/%s?recursive=1", github["reposApi"], github["repo"], github["branch"])
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	repoTreeResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var repoTree tree
	err = json.Unmarshal([]byte(repoTreeResponse), &repoTree)
	if err != nil {
		panic(err)
	}

	var fileList []string
	for _, v := range repoTree.Tree {
		if strings.Contains(v.Path, ".gitignore") {
			fileList = append(fileList, strings.ReplaceAll(v.Path, ".gitignore", ""))
		}
	}

	return fileList
}

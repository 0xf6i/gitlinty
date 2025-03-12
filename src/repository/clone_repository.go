package repository

import (
	"errors"
	"fmt"
	"linty/src/config"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

func Clone(author string, repo string, config *config.Config) (string, string, error) {

	// Generate a unique folder name
	folderName := author + "-" + repo
	fmt.Println("generated unique folder:", folderName)

	// Construct the full clone path using the configured directory
	clonePath := filepath.Join(config.DirectoryPath, "gitlinty", folderName)
	repoURL := "https://github.com/" + author + "/" + repo
	fmt.Println("cloning:", repoURL, "into", clonePath)

	if _, err := os.Stat(clonePath); err == nil {
		err := os.RemoveAll(clonePath)
		if err != nil {
			return "", "", errors.New("failed to remove already existing repository")
		} else {
			fmt.Println("removed current directory", clonePath)
		}

	}

	_, err := git.PlainClone(clonePath, false, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		return "", "", errors.New("failed to clone repository")
	}

	return clonePath, folderName, nil
}

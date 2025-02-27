package repository

import (
	"fmt"
	"linty/utils"

	"github.com/go-git/go-git/v5"
)

func Clone(author string, repo string) (string, string, error) {
	config, err := utils.LoadConfig("config.json")
	if err != nil {
		return "", "", fmt.Errorf("failed to load config: %w", err)
	}

	// If config is missing, return an error
	if config == nil {
		return "", "", fmt.Errorf("config file not found or invalid")
	}

	// Generate a unique folder name
	base64 := utils.GenerateBase64()
	fmt.Println("Generated Unique Folder:", base64)

	// Construct the full clone path using the configured directory
	clonePath := config.DirectoryPath + "/" + base64
	repoURL := "https://github.com/" + author + "/" + repo
	fmt.Println("Cloning:", repoURL, "into", clonePath)

	// Clone the repository
	_, err = git.PlainClone(clonePath, false, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		return "", "", fmt.Errorf("failed to clone repo: %w", err)
	}

	return clonePath, base64, nil
}

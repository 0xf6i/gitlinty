package repository

import (
	"errors"
	"os"
	"path/filepath"
)

// CheckRepository takes a path and checks if the given path is a directory and if this said directory contains a .git folder inside.
func CheckRepository(path string) (string, error) {
	path = filepath.Clean(path)
	dirStat, err := os.Stat(path)
	if err != nil {
		return "", errors.New("failed to access path:" + path)
	}
	if dirStat.IsDir() {
		gitPath := filepath.Join(path, ".git")

		gitStat, err := os.Stat(gitPath)
		if err != nil {
			if os.IsNotExist(err) {
				return "", errors.New(".git does not exist in given directory: " + gitPath)
			}
			return "", errors.New("Failed to access .git directory: " + gitPath)
		}
		if gitStat.IsDir() {
			return gitPath, nil
		}
	}
	return "", errors.New(".git exists but is not in a directory")
}

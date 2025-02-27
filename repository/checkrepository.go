package repository

import (
	"errors"
	"os"
)

// CheckRepository takes a path and checks if the given path is a directory and if this said directory contains a .git folder inside.
func CheckRepository(path string) (string, bool, bool, error) {
	dirStat, err := os.Stat(path)
	if err != nil {
		return "", false, false, errors.New("Failed to access path:" + path)
	}
	if dirStat.IsDir() {
		gitPath := path + "/.git"

		gitStat, err := os.Stat(gitPath)
		if err != nil {
			if os.IsNotExist(err) {
				return "", true, false, errors.New(".git does not exist in given directory: " + gitPath)
			}
			return "", true, false, errors.New("Failed to access .git directory: " + gitPath)
		}
		if gitStat.IsDir() {
			return gitPath, true, true, nil
		}
	}
	return "", true, false, errors.New(".git exists but is not in a directory")
}

package local

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// CheckRepository takes a path and checks if the given path is a directory and if this said directory contains a .git folder inside.
func CheckRepository(path string) (string, bool, bool, error) {
	dirStat, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	if dirStat.IsDir() {
		fmt.Println("[checkrepository.go]: directory exists.")
		gitPath := path + "/.git"

		gitStat, err := os.Stat(gitPath)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("[checkrepository.go]: gitPath: " + gitPath)

		if gitStat.IsDir() {
			fmt.Println("[checkrepository.go]: .git directory exists.")
			return gitPath, true, true, nil
		}
		return "", true, false, errors.New("[checkrepository.go]: .git does not exist")

	}
	return "", false, false, errors.New("[checkrepository.go]: neither dir or git exists")

}

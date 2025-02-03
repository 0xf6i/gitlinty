package local

import (
	"errors"
	"fmt"
	"os"
)

func Fetch(path string) (bool, error) {
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		fmt.Println("directory exists.")
		p := path + "/.git"
		if _, err := os.Stat(p); err == nil {
			return true, nil
		}
		return false, errors.New(".git directory does not exist")
	}
	return false, errors.New("directory does not exist")
}

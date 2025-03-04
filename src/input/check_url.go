package input

import (
	"errors"
	"linty/src/url"
	"os"
	"strings"
)

func CheckUrl(path string) (bool, error) {
	path = strings.ToLower(path)
	if strings.Contains(path, "http") || strings.Contains(path, "https") || strings.Contains(path, "github.com") {
		url, err := url.CheckValidity(path)
		if err != nil {
			return false, errors.New("error checking validity of url, please make sure to enter the entire url")
		}
		if url {
			return true, nil
		}
	} else {
		dir, err := os.Stat(path)
		if err != nil {
			return false, errors.New("given directory does not seem to be an existing directory")
		}
		if dir.IsDir() {
			return false, nil
		}
	}
	return false, errors.New("invalid path")
}

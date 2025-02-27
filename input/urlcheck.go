package input

import (
	"errors"
	"linty/url"
	"log"
	"os"
	"strings"
)

func CheckIfUrl(path string) (bool, error) {
	if strings.Contains(path, "http") || strings.Contains(path, "https") || strings.Contains(path, "github.com") {
		url, _ := url.CheckValidity(path)
		if url {
			return true, nil
		}
	} else {
		dir, err := os.Stat(path)
		if err != nil {
			log.Fatal(err)
		}

		if dir.IsDir() {
			return false, nil
		}
	}
	return false, errors.New("invalid path")
}

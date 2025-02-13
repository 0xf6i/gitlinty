package url

import (
	"errors"
	"strings"
)

type HandlerStruct struct {
	Author     string
	Repository string
}

func Handler(url string) (HandlerStruct, error) {
	splitString := strings.Split(url, "github.com/")
	if len(splitString) < 2 {
		return HandlerStruct{}, errors.New("INVALID URL")
	}
	parts := strings.Split(splitString[1], "/")
	if len(parts) < 2 || parts[0] == "" || parts[1] == "" {
		return HandlerStruct{}, errors.New(" URL IS MISSING AUTHOR OR REPOSITORY")
	}

	return HandlerStruct{
		Author:     parts[0],
		Repository: parts[1],
	}, nil
}

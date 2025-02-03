package urlhandler

import (
	"errors"
	"fmt"
	"strings"
)

type UrlHandlerStruct struct {
	Author     string
	Repository string
}

func UrlHandler(url string) (UrlHandlerStruct, error) {
	fmt.Println("[URL HANDLER]: SPLITTING STRING")
	splitString := strings.Split(url, "github.com/")
	if len(splitString) < 2 {
		return UrlHandlerStruct{}, errors.New("INVALID URL")
	}
	parts := strings.Split(splitString[1], "/")
	if len(parts) < 2 || parts[0] == "" || parts[1] == "" {
		return UrlHandlerStruct{}, errors.New("URL IS MISSING AUTHOR OR REPOSITORY")
	}
	fmt.Println("[URL HANDLER]: STRING SPLIT SUCCESSFULLY")

	fmt.Println("[URL HANDLER]: RETURNING STRUCT")
	return UrlHandlerStruct{
		Author:     parts[0],
		Repository: parts[1],
	}, nil
}

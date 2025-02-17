package repository

import (
	"fmt"
	"linty/utils"

	"github.com/go-git/go-git/v5"
)

func Clone(author string, repo string) (string, error) {
	uuid := utils.GenerateUuid()
	fmt.Println(uuid)
	url := "https://github.com/" + author + "/" + repo
	fmt.Println(url)

	_, err := git.PlainClone("/tmp/gitlinty/"+uuid, false, &git.CloneOptions{
		URL: url,
	})

	if err != nil {
		return "", err
	}

	returnPath := "/tmp/gitlinty/" + uuid

	return returnPath, nil

}

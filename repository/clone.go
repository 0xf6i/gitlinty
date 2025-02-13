package repository

import (
	"linty/utils"

	"github.com/go-git/go-git/v5"
)

func Clone(author string, repo string) (string, error) {
	uuid := utils.GenerateUuid()

	_, err := git.PlainClone("/tmp/gitlinty/"+uuid, false, &git.CloneOptions{
		URL: "https://github.com/" + author + "/" + repo,
	})

	if err != nil {
		return "", err
	}

	returnPath := "/tmp/gitlinty/" + uuid

	return returnPath, nil

}

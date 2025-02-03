package repository

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"linty/utils"
)

func Clone(author string, repo string) (string, error) {
	fmt.Println("[UTILS]: GENERATING UUID")
	uuid := utils.GenerateUuid()
	fmt.Println("[UTILS]: GENERATED UUID")

	_, err := git.PlainClone("/tmp/"+uuid, false, &git.CloneOptions{
		URL: "https://github.com/" + author + "/" + repo,
	})

	if err != nil {
		return "", err
	}
	fmt.Println("[REPOSITORY]: CLONED REPOSITORY SUCCESSFULLY")

	fmt.Println("[REPOSITORY]: GENERATING TEMP PATH")
	returnPath := "/tmp/" + uuid

	fmt.Println("[REPOSITORY]: RETURNING TEMP PATH")
	return returnPath, nil

}

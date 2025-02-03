package main

import (
	"errors"
	"fmt"
	git "github.com/go-git/go-git/v5"
	"linty/input"
	"linty/local"
	"linty/repository"
	"linty/urlhandler"
	"linty/utils"
	"log"
	"os"
	"strings"
)

func cloneRespository(path string) {
	id := utils.GenerateUuid()
	fmt.Println(id)
	_, err := git.PlainClone("/tmp/"+id, false, &git.CloneOptions{
		URL:      path,
		Progress: nil,
	})
	if err != nil {
		log.Fatalf("Failed to clone repository: %v", err)
	}
	fmt.Println("Repository cloned successfully!")
	dir := "/tmp/" + id + "/.git"
	if _, err := os.Stat(dir); err == nil {
		fmt.Println("Directory exists")
		fmt.Println(string(dir))
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println(dir + " does not exists")
	} else {
		fmt.Println("unsure if " + dir + " exists")
	}

}

func main() {
	s, e := input.InputReader("Please specify a route:")
	if e != nil {
		errors.New("")
	}
	prefixTrueFalse := strings.HasPrefix(s, "http")
	fmt.Println("[MAIN]: CHECKING FOR URL OR LOCAL")

	switch prefixTrueFalse {
	case true:
		fmt.Println("[MAIN]: GOT URL AS INPUT")
		url := s
		fmt.Println("[URL HANDLER]: CHECKING URL VALIDITY")
		cuv, e := urlhandler.CheckUrlValidity(url)
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println("[URL HANDLER]: URL VALID")
		if cuv == true {
			uhs, e := urlhandler.UrlHandler(url)
			if e != nil {
				fmt.Println(e)
			}
			fmt.Println("[URL HANDLER]: URL HANDLED")

			fmt.Println("[REPOSITORY]: TRYING TO CLONE REPOSITORY")
			uuid, e := repository.Clone(uhs.Author, uhs.Repository)
			if e != nil {
				fmt.Println(e)
			}
			fmt.Println("[MAIN]: CLONED REPOSITORY TO " + uuid)

		}
	default:
		log.Println("path")
		l, e := local.Fetch(s)
		if e != nil {
			errors.New("error")
		}
		fmt.Println(l)
	}

}

package main

import (
	"fmt"
	"linty/input"
	"linty/local"
	"linty/repository"
	"linty/url"
	"linty/utils"
	"log"
	"runtime"

	"github.com/fatih/color"
)

var version = "0.1.2"

//func main() {
//	os := runtime.GOOS
//	color.Blue("Gitlinty for " + os + " | " + version + " (https://github.com/0xf6i/gitlinty/)")
//
//	userUrlInput, err := input.UserInput("PATH")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("url: " + userUrlInput)
//
//	handledUrl, err := url.Handler(userUrlInput)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(handledUrl.Author)
//	fmt.Println(handledUrl.Repository)
//
//	clonedRepository, base64, err := repository.Clone(handledUrl.Author, handledUrl.Repository)
//	if err != nil {
//		fmt.Println("err")
//		log.Fatal(err)
//	}
//	fmt.Println(clonedRepository)
//
//	contributors, err := repository.CheckContributors(clonedRepository)
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//
//	noOfCommits := summary.SummarizeCommits(contributors)
//	fmt.Println(noOfCommits)
//
//	fmt.Println(utils.DecodeBase64(base64))
//
//}

func main() {
	defaultPath := "/tmp/gitlinty"
	os := runtime.GOOS
	fmt.Println("✅")
	color.Blue("Gitlinty for " + os + " (DEMO) | " + version + " (https://github.com/0xf6i/gitlinty/)")

	success, directory := utils.CheckIfGitlintyExists(defaultPath)
	if success {
		fmt.Println("Gitlinty is ready at:", directory)
		pathInput, err := input.UserInput("Please specify an URL or a local path which you want to validate using Gitlinty:")
		if err != nil {
			log.Fatal(err)
		}

		urlorpath, err := input.CheckIfUrl(pathInput)
		if err != nil {
			log.Fatal(err)
		}

		switch urlorpath {
		case true:
			urlvalid, err := url.CheckValidity(pathInput)
			if err != nil {
				fmt.Println(err)
			}
			if urlvalid {
				rep, _ := url.Handler(pathInput)
				clonePath, _, err := repository.Clone(rep.Author, rep.Repository)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Successfully cloned repository to: " + clonePath)

			}
		case false:
			gitPath, dir, git, err := local.CheckRepository(pathInput)
			if err != nil {
				fmt.Println(err)
			}
			color.Blue(".git exists in given directory: %t at %s", git, gitPath)
			color.Blue("directory exists: %t", dir)
		}
	} else {
		fmt.Println("Failed to set up Gitlinty.")
	}
}

//urlValidity, err := url.CheckValidity(pathInput)
//if err != nil {
//	log.Fatal(err)
//}
//fmt.Println("[GITLINTY] Url valid: ")
//color.Green("%t", urlValidity)
//
//if urlValidity {
//
//	handledUrl, err := url.Handler(pathInput)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("[GITLINTY]: Parsed repository author: " + handledUrl.Author)
//	fmt.Println("[GITLINTY]: Parsed repository name: " + handledUrl.Repository)
//} else {
//}

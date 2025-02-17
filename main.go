package main

import (
	"fmt"
	"linty/input"
	"linty/repository"
	"linty/summary"
	"linty/url"
	"log"
	"runtime"

	"github.com/fatih/color"
)

var version = "0.1.2"

func main() {
	os := runtime.GOOS
	color.Blue("Gitlinty for " + os + " | " + version + " (https://github.com/0xf6i/gitlinty/)")

	userUrlInput, err := input.UserInput("PATH")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("url: " + userUrlInput)

	handledUrl, err := url.Handler(userUrlInput)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(handledUrl.Author)
	fmt.Println(handledUrl.Repository)

	clonedRepository, err := repository.Clone(handledUrl.Author, handledUrl.Repository)
	if err != nil {
		fmt.Println("err")
		log.Fatal(err)

	}
	fmt.Println(clonedRepository)

	contributors, err := repository.CheckContributors(clonedRepository)
	if err != nil {
		log.Fatal(err)
		return
	}

	noOfCommits := summary.SummarizeCommits(contributors)
	fmt.Println(noOfCommits)

}

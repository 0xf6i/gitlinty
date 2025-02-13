package main

import (
	"fmt"
	"github.com/fatih/color"
	"linty/input"
	"linty/local"
	"log"
	"runtime"
)

var version = "0.1.1"

func main() {
	os := runtime.GOOS
	color.Blue("Gitlinty " + version + " (https://github.com/0xf6i/gitlinty/)")
	fmt.Println(os)

	filePath, err := input.UserInput("Please give me text file path")
	if err != nil {
		log.Fatal(err)
	}
	fileContainsTest, err := local.CheckForTest(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileContainsTest)

}

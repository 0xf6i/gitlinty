package main

import (
	"fmt"
	"linty/src/config"
	"linty/src/files"
	"linty/src/input"
	"linty/src/repository"
	"linty/src/url"
	"linty/src/utils"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/fatih/color"
)

//if default path -> ask for input
//if custom path -> check path exists -> create directory -> save to config

// ask for input -> url -> validate -> handle -> clone -> run checks
// ask for input -> path -> validate -> run checks

// return summary object -> print summary

func main() {
	opersys := runtime.GOOS
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("failed to read current directory")
	}
	// TESTING
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("failed to get user home directory")
	}

	color.Blue("Gitlinty for " + opersys + " (DEMO) |  (https://github.com/0xf6i/gitlinty/)")

	// load config
	loadedConfig, err := config.LoadConfig(filepath.Join(currentDir, "config.json"))
	if err != nil {
		log.Fatal("failed to read current config file")
	}
	fmt.Println("Config file has been loaded.")

	//clean clone directory path
	cloneDirectory := filepath.Clean(loadedConfig.DirectoryPath)

	// check if given directory in config exists
	cloneDirectoryExists := utils.FolderExists(cloneDirectory)
	if err != nil {
		log.Fatal(err)
	}
	// if the directory doesnt exist, create it
	if !cloneDirectoryExists {
		os.Mkdir(filepath.Join(cloneDirectory, "gitlinty"), os.ModePerm)
	}

	//HARDCODED TESTING VALUE, REMOVE LATER
	// ß

	// run the first time sequence
	if loadedConfig.FirstRun {
		// check if user wants custom path of default one
		useDefaultPath, err := input.UserChoice("Would you like to use the default path?")
		if err != nil {
			log.Fatal(err)
		}
		// user wants custom path
		if !useDefaultPath {
			userSelectedPath, err := input.UserInput("Please specify the path to where you want to stored cloned repositories")
			if err != nil {
				log.Fatal(err)
			}
			// check if the given path and directory exists
			userSelectedPathExists := utils.FolderExists(userSelectedPath)
			if err != nil {
				log.Fatal(err)
			}
			// path does not exist, ask if they want to create it, otherwise exit
			if !userSelectedPathExists {
				userCreateNewDirectory, err := input.UserChoice("Directory does not exist, do you want to create a new directory?")
				if err != nil {
					log.Fatal(err)
				}
				if userCreateNewDirectory {
					os.MkdirAll(userSelectedPath, os.ModePerm)
				} else {
					fmt.Println("Exiting program...")
					os.Exit(0)
				}
			}
			//clean given directory, set firstRun = false and save config
			loadedConfig.DirectoryPath = filepath.Clean(userSelectedPath)
			loadedConfig.FirstRun = false
			config.WriteConfig(loadedConfig, filepath.Join(currentDir, "config.json"))
		} else {
			//clean default directory, set firstRun = false and save config
			loadedConfig.DirectoryPath = filepath.Join(userHomeDir, "tmp", "gitlinty")
			loadedConfig.FirstRun = false
			config.WriteConfig(loadedConfig, filepath.Join(currentDir, "config.json"))
		}
	}
	//reload config file
	loadedConfig, err = config.LoadConfig(filepath.Join(currentDir, "config.json"))
	if err != nil {
		log.Fatal()
	}
	fmt.Println("Program will clone repositories to:", loadedConfig.DirectoryPath)

	userClonePath, err := input.UserInput("Please specify the path to the repository you want to clone:")
	if err != nil {
		log.Fatal(err)
	}

	pathIsUrl, err := input.CheckUrl(userClonePath)
	if err != nil {
		log.Fatal(err)
	}
	switch pathIsUrl {
	case true:
		validUrl, err := url.CheckValidity(userClonePath)
		if err != nil {
			log.Fatal(err)
		}
		if validUrl {
			handledUrl, err := url.Handler(userClonePath)
			if err != nil {
				log.Fatal(err)
			}
			repo := handledUrl.Repository
			author := handledUrl.Author

			clonedRepoPath, directoryName, err := repository.Clone(author, repo, loadedConfig)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(clonedRepoPath)
			fmt.Println(directoryName)

			licenseFiles, err := files.CheckFileContent(clonedRepoPath, "license")
			if err != nil {
				log.Fatal(err)
			}
			gitIgnoreFiles, err := files.CheckFileContent(clonedRepoPath, "gitignore")
			if err != nil {
				log.Fatal(err)
			}
			readmeFiles, err := files.CheckFileContent(clonedRepoPath, "readme")
			if err != nil {
				log.Fatal(err)
			}
			workFlowFiles, err := files.CheckFileContent(clonedRepoPath, "workflow")
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(licenseFiles, gitIgnoreFiles, readmeFiles, workFlowFiles)
		}

	case false:
		fmt.Println("PATH")
	}

	// loadedConfig.DirectoryPath = filepath.Join(currentDir, "folder_to_clone_to")

	// newConfig, err := config.WriteConfig(loadedConfig, filepath.Join(currentDir, "config.json"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(newConfig)

}

package main

import (
	"fmt"
	"github.com/fatih/color"
	"linty/src/config"
	"linty/src/execute"
	"linty/src/input"
	"linty/src/utils"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func main() {

	opersys := runtime.GOOS
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println("failed to read current directory")
		os.Exit(1)

	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("failed to get user home directory")
		os.Exit(1)
	}

	color.Blue("Gitlinty for " + opersys + " (DEMO) |  (https://github.com/0xf6i/gitlinty/)")

	loadedConfig, err := config.LoadConfig(filepath.Join(currentDir, "config.json"))
	if err != nil {
		log.Println("failed to read current config file")
		os.Exit(1)
	}
	fmt.Println("Config file has been loaded.")

	cloneDirectory := filepath.Clean(loadedConfig.DirectoryPath)

	cloneDirectoryExists := utils.FolderExists(cloneDirectory)
	if err != nil {
		log.Println(err)
	}

	if !cloneDirectoryExists {
		os.Mkdir(filepath.Join(cloneDirectory, "gitlinty"), os.ModePerm)
	}

	if loadedConfig.FirstRun {

		useDefaultPath, err := input.UserChoice("Would you like to use the default path?")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		if !useDefaultPath {
			userSelectedPath, err := input.UserInput("Please specify the path to where you want to stored cloned repositories")
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}

			userSelectedPathExists := utils.FolderExists(userSelectedPath)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}

			if !userSelectedPathExists {
				userCreateNewDirectory, err := input.UserChoice("Directory does not exist, do you want to create a new directory?")
				if err != nil {
					log.Println(err)
					os.Exit(1)
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

	loadedConfig, err = config.LoadConfig(filepath.Join(currentDir, "config.json"))
	if err != nil {
		log.Println()
	}
	fmt.Println("Program will clone repositories to:", loadedConfig.DirectoryPath)
	userClonePath, err := input.UserInput("Please specify the path to the repository you want to clone:")
	execute.Execute(userClonePath, loadedConfig)

}

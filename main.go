package main

import (
	"fmt"
	"linty/src/config"
	"linty/src/files"
	"linty/src/input"
	"linty/src/repository"
	"linty/src/summary"
	"linty/src/url"
	"linty/src/utils"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

func main() {

	opersys := runtime.GOOS
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("failed to read current directory")
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("failed to get user home directory")
	}

	color.Blue("Gitlinty for " + opersys + " (DEMO) |  (https://github.com/0xf6i/gitlinty/)")

	loadedConfig, err := config.LoadConfig(filepath.Join(currentDir, "config.json"))
	if err != nil {
		log.Fatal("failed to read current config file")
	}
	fmt.Println("Config file has been loaded.")

	cloneDirectory := filepath.Clean(loadedConfig.DirectoryPath)

	cloneDirectoryExists := utils.FolderExists(cloneDirectory)
	if err != nil {
		log.Fatal(err)
	}

	if !cloneDirectoryExists {
		os.Mkdir(filepath.Join(cloneDirectory, "gitlinty"), os.ModePerm)
	}

	if loadedConfig.FirstRun {

		useDefaultPath, err := input.UserChoice("Would you like to use the default path?")
		if err != nil {
			log.Fatal(err)
		}

		if !useDefaultPath {
			userSelectedPath, err := input.UserInput("Please specify the path to where you want to stored cloned repositories")
			if err != nil {
				log.Fatal(err)
			}

			userSelectedPathExists := utils.FolderExists(userSelectedPath)
			if err != nil {
				log.Fatal(err)
			}

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

			fmt.Println("CLONED REPO PATH:", clonedRepoPath)
			fmt.Println("DIRECTORY NAME", directoryName)

			ignoredPaths, ignoredPatterns, err := summary.ReadGitignore(clonedRepoPath)
			if err != nil {
				log.Printf("Warning: couldn't read .gitignore: %v", err)
				ignoredPaths = []string{}
				ignoredPatterns = []string{}
			}

			allFiles, err := files.FindProjectFiles(clonedRepoPath, ignoredPaths, ignoredPatterns)
			if err != nil {
				log.Fatal(err)
			}

			repoObject := summary.Repository{
				Author: handledUrl.Author,
				Name:   handledUrl.Repository,
			}

			categories := []string{"license", "gitignore", "readme", "workflow", "tests"}
			//fmt.Println("CLoned repo path:", clonedRepoPath)
			results := summary.GenerateSummary(repoObject, allFiles, categories, clonedRepoPath, loadedConfig)

			summary.PrintSummary(results)

			cmd := exec.Command("gitleaks", "git", "-v", clonedRepoPath)
			output, err := cmd.CombinedOutput()

			if err != nil {
				if strings.Contains(string(output), "leaks found") {
					fmt.Println("\n⚠️ Gitleaks found potential secrets in the repository:")
					fmt.Println(string(output))
					os.Exit(1)
				}
				log.Fatal(err)
			}

			fmt.Println("\nGitleaks scan completed successfully - no leaks found")
			os.Exit(0)

		}

	case false:
		ignoredPaths, ignoredPatterns, err := summary.ReadGitignore(userClonePath)
		if err != nil {
			log.Printf("Warning: couldn't read .gitignore: %v", err)
			ignoredPaths = []string{}
			ignoredPatterns = []string{}
		}

		allFiles, err := files.FindProjectFiles(userClonePath, ignoredPaths, ignoredPatterns)
		if err != nil {
			log.Fatal(err)
		}

		author, repo, err := repository.GetRepoInfo(userClonePath)
		if err != nil {
			log.Fatal(err)
		}

		repoObject := summary.Repository{
			Author: author,
			Name:   repo,
		}
		categories := []string{"license", "gitignore", "readme", "workflow", "tests"}
		results := summary.GenerateSummary(repoObject, allFiles, categories, userClonePath, loadedConfig)

		summary.PrintSummary(results)
		cmd := exec.Command("gitleaks", "git", "-v", userClonePath)
		output, err := cmd.CombinedOutput()

		if err != nil {
			if strings.Contains(string(output), "leaks found") {
				fmt.Println("\n⚠️ Gitleaks found potential secrets in the repository:")
				fmt.Println(string(output))
				os.Exit(0)
			}
			log.Fatal(err)
		}

		fmt.Println("\nGitleaks scan completed successfully - no leaks found")
		fmt.Println("Exiting with status code: 0")
		os.Exit(0)

	}

}

package execute

import (
	"fmt"
	"linty/src/config"
	"linty/src/files"
	"linty/src/input"
	"linty/src/repository"
	"linty/src/summary"
	"linty/src/url"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Execute(userInput string, loadedConfig *config.Config) {
	userClonePath := userInput

	pathIsUrl, err := input.CheckUrl(userClonePath)
	if err != nil {
		log.Fatal("couldn't not check url: ", err)

	}
	switch pathIsUrl {
	case true:
		validUrl, err := url.CheckValidity(userClonePath)
		if err != nil {
			log.Fatal("couldn't not validate url: ", err)

		}
		if validUrl {
			handledUrl, err := url.Handler(userClonePath)
			if err != nil {
				log.Fatal("couldn't not handle url: ", err)

			}
			repo := handledUrl.Repository
			author := handledUrl.Author

			clonedRepoPath, _, err := repository.Clone(author, repo, loadedConfig)
			if err != nil {
				log.Fatal("couldn't not clone repository: ", err)
			}

			ignoredPaths, ignoredPatterns, err := summary.ReadGitignore(clonedRepoPath)
			if err != nil {
				log.Fatal("couldn't read .gitignore: ", err)
			}

			allFiles, err := files.FindProjectFiles(clonedRepoPath, ignoredPaths, ignoredPatterns)
			if err != nil {
				log.Fatal("couldn't not read project files: ", err)

			}

			repoObject := summary.Repository{
				Author: handledUrl.Author,
				Name:   handledUrl.Repository,
			}

			categories := []string{"license", "gitignore", "readme", "workflow", "tests"}
			results := summary.GenerateSummary(repoObject, allFiles, categories, clonedRepoPath, loadedConfig)

			summary.PrintSummary(results)

			cmd := exec.Command("gitleaks", "git", "-v", clonedRepoPath)
			output, err := cmd.CombinedOutput()

			if err != nil {
				if strings.Contains(string(output), "leaks found") {
					fmt.Println("\n⚠️ gitleaks found potential secrets in the repository:")
					fmt.Println(string(output))
				}
				//log.Fatal("gitleaks found errors: ", err)
			}

			fmt.Println("\ngitleaks scan completed successfully - no leaks found")

			removeFolder, err := input.UserChoice("Do you want to remove the cloned repository?")
			if err != nil {
				log.Fatal("couldn't not read user input: ", err)
			}
			if removeFolder {
				os.RemoveAll(clonedRepoPath)
			}

			fmt.Println("Exiting with status code: 0")
			os.Exit(0)
		}

	case false:
		ignoredPaths, ignoredPatterns, err := summary.ReadGitignore(userClonePath)
		if err != nil {
			log.Fatal("couldn't read .gitignore: ", err)

		}

		allFiles, err := files.FindProjectFiles(userClonePath, ignoredPaths, ignoredPatterns)
		if err != nil {
			log.Fatal("couldn't not read project files: ", err)
		}

		author, repo, err := repository.GetRepoInfo(userClonePath)
		if err != nil {
			log.Fatal("couldn't not get repository information: ", err)
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
				fmt.Println("\n⚠️ gitleaks found potential secrets in the repository:")
				fmt.Println(string(output))
			}
			//log.Fatal("gitleaks found errors: ", err)
		}

		fmt.Println("\ngitleaks scan completed successfully - no leaks found")

		removeFolder, err := input.UserChoice("Do you want to remove the cloned repository?")
		if err != nil {
			log.Fatal("couldn't not read user input: ", err)
		}
		if removeFolder {
			os.RemoveAll(userClonePath)
		}

		fmt.Println("Exiting with status code: 0")
		os.Exit(0)

	}
}

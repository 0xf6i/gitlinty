package repository

import (
	"fmt"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Contributor struct {
	Name    string
	Email   string
	Commits int
}

func CheckContributors(path string) (*[]Contributor, error) {
	path = filepath.Clean(path)
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open repository: %w", err)
	}

	// Get the HEAD reference
	ref, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("failed to get HEAD: %w", err)
	}

	// Get the commit history
	commitIteration, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return nil, fmt.Errorf("failed to get commit history: %w", err)
	}
	summary := make(map[string]*Contributor)

	err = commitIteration.ForEach(func(c *object.Commit) error {
		key := c.Author.Name
		if _, exists := summary[key]; !exists {
			summary[key] = &Contributor{
				Name:    c.Author.Name,
				Email:   c.Author.Email,
				Commits: 0,
			}
		}

		summary[key].Commits++
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error iterating commits: %w", err)
	}

	var contributors []Contributor
	for _, contributor := range summary {

		contributors = append(contributors, *contributor)
	}

	sort.Slice(contributors, func(i, j int) bool {
		return contributors[i].Commits > contributors[j].Commits
	})

	return &contributors, nil

}

// GetRepoInfo extracts the repository name and author (username) from a local Git repository
func GetRepoInfo(path string) (string, string, error) {
	// Open the repository at the given path
	repo, err := git.PlainOpen(path)
	if err != nil {
		return "", "", fmt.Errorf("failed to open repository: %w", err)
	}

	// Try to get remote origin (for remote repositories)
	remote, err := repo.Remote("origin")
	if err == nil && len(remote.Config().URLs) > 0 {
		remoteURL := remote.Config().URLs[0]

		// Normalize SSH URLs to HTTPS format
		remoteURL = strings.Replace(remoteURL, "git@", "https://", 1)
		remoteURL = strings.Replace(remoteURL, ":", "/", 1)

		// Regex to extract username and repository name
		re := regexp.MustCompile(`.*/([^/]+)/([^/]+)\.git`)
		matches := re.FindStringSubmatch(remoteURL)
		if len(matches) == 3 {
			return matches[1], matches[2], nil
		}
	}

	// If no remote, fallback to local folder name as repo name
	repoName := extractRepoNameFromPath(path)

	// Get first commit author as repo owner
	author, err := getFirstCommitAuthor(repo)
	if err != nil {
		return "", "", fmt.Errorf("failed to determine repository owner: %w", err)
	}

	return author, repoName, nil
}

// extractRepoNameFromPath gets the repository name from the folder name
func extractRepoNameFromPath(path string) string {
	parts := strings.Split(strings.TrimRight(path, "/"), "/")
	if len(parts) == 0 {
		return "unknown-repo"
	}
	return parts[len(parts)-1]
}

// getFirstCommitAuthor retrieves the first commit author's name
func getFirstCommitAuthor(repo *git.Repository) (string, error) {
	ref, err := repo.Head()
	if err != nil {
		return "", fmt.Errorf("failed to get HEAD: %w", err)
	}

	commitIter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return "", fmt.Errorf("failed to get commit history: %w", err)
	}

	lastCommit, err := commitIter.Next()
	if err != nil {
		return "", fmt.Errorf("failed to get first commit: %w", err)
	}

	return lastCommit.Author.Name, nil
}

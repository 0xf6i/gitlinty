package repository

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Contributor struct {
	Name    string
	Email   string
	Commits int
}

func CheckContributors(path string) (*[]Contributor, error) {
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

	return &contributors, nil

}

package summary

import (
	"fmt"
	"linty/src/repository"
	"sort"
)

func SummarizeCommits(contributors *[]repository.Contributor) int {
	sort.SliceStable(*contributors, func(i, j int) bool {
		return (*contributors)[i].Commits > (*contributors)[j].Commits
	})

	totalCommits := 0

	for _, contributor := range *contributors {
		totalCommits += contributor.Commits
		fmt.Printf("Name: %s, Email: %s, Commits: %d\n", contributor.Name, contributor.Email, contributor.Commits)
	}
	return totalCommits
}

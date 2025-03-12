package summary

func SummarizeCommits(contributors []Contributor) int {
	totalCommits := 0

	// Iterate over contributors and display their details
	for _, contributor := range contributors {
		totalCommits += contributor.Commits
	}

	return totalCommits
}

package summary

import (
	"fmt"
	"linty/src/config"
	"linty/src/repository"
)

// GenerateSummary creates a summary of the repository analysis
func GenerateSummary(repo Repository, allFiles [][]File, categories []string, clonedRepoPath string, config *config.Config) *SummaryResult {
	summary := Summary{
		Repository: repo,
	}

	GetFailureAllowance(config, &summary)

	redLight := false
	categoryResults := make(map[string][]File)
	status := make(map[string]string) // Tracks traffic light status for each category
	reason := make(map[string]string) // Tracks the reason for red/yellow light

	// Create a map of files to skip for faster lookup
	filesToSkipMap := make(map[string]bool)
	for _, skip := range config.FilesToSkip {
		if path, ok := skip.(string); ok {
			filesToSkipMap[path] = true
			fmt.Printf("Due to config program will skip file: %s\n", path)
		}
	}

	for i, category := range categories {
		files := allFiles[i]
		var passed, warning, failed []File

		for _, file := range files {
			// Check if this file is in the skip list from config
			if filesToSkipMap[file.Path] {

				// Even if file is skipped, we still need to track its status for reporting
				// but we won't count it against the category status
				if file.Note == "Empty file" {
					file.Status = "yellow" // Add status but don't affect category

					// Add to warnings so it shows up in the report, but marked as skipped
					file.Note = file.Note + " (skipped by config)"

					if _, exists := reason[category]; !exists {
						reason[category] = "File has issues but is skipped by config"
					}
				} else {
					// If it's a skipped file with no issues, we can still mark it as green
					file.Status = "green"
					file.Note = "Skipped by config"
					passed = append(passed, file)
				}
				continue
			}

			if file.Note == "Empty file" {
				file.Status = "yellow"
				warning = append(warning, file)
				if _, exists := reason[category]; !exists {
					reason[category] = "File is empty or contains only whitespace"
				}
			} else {
				file.Status = "green"
				passed = append(passed, file)
			}
		}

		// Handle missing files
		if len(files) == 0 {
			if IsFailureAllowed(category, summary) {
				missingFile := File{
					Path:   "N/A",
					Type:   category,
					Note:   "File is missing but allowed to fail",
					Status: "yellow",
				}
				warning = append(warning, missingFile)
				reason[category] = "File is missing but allowed to fail"
			} else {
				if category == "tests" || category == "workflow" {
					continue
				}
				missingFile := File{
					Path:   "N/A",
					Type:   category,
					Note:   "Missing file",
					Status: "red",
				}
				failed = append(failed, missingFile)
				reason[category] = "File is missing"
			}
		}

		// Determine status
		if len(failed) > 0 {
			status[category] = "red"
			redLight = true
		} else if len(warning) > 0 {
			status[category] = "yellow"
			if _, exists := reason[category]; !exists {
				reason[category] = "File has warnings"
			}
		} else if len(passed) > 0 {
			status[category] = "green"
		}

		// Add results to category - always include the failed files
		// Even if we're stopping early due to redLight
		// if len(failed) > 0 {

		// }

		// Only add passed and warning files if not in redLight mode
		if !redLight {
			categoryResults[category] = append(categoryResults[category], passed...)
			categoryResults[category] = append(categoryResults[category], warning...)
		}

		if redLight {
			categoryResults[category] = append(categoryResults[category], failed...)
		}
	}

	contributors, err := repository.CheckContributors(clonedRepoPath)

	if err != nil {
		fmt.Println("Error fetching contributors:", err)
	}
	totalCommits := 0
	var summaryContributors []Contributor

	for _, repoContributor := range *contributors {

		summaryContributors = append(summaryContributors, Contributor{
			Name:    repoContributor.Name,
			Email:   repoContributor.Email,
			Commits: repoContributor.Commits,
		})
		totalCommits += repoContributor.Commits
	}

	summary.Contributions = Contributions{
		TotalCommits: totalCommits,
		Contributors: summaryContributors,
	}

	return &SummaryResult{
		Summary:         &summary,
		CategoryResults: categoryResults,
		Status:          status,
		Reason:          reason,
		RedLight:        redLight,
	}
}

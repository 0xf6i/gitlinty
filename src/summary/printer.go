package summary

import (
	"fmt"
	"linty/src/input"
	"log"
	"os"
	"strings"
)

func PrintSummary(result *SummaryResult) {
	// Create maps to organize files by status
	greenFiles := make(map[string][]File)
	yellowFiles := make(map[string][]File)
	redFiles := make(map[string][]File)

	// Sort files by their status
	for category, files := range result.CategoryResults {
		for _, file := range files {
			switch file.Status {
			case "green":
				greenFiles[category] = append(greenFiles[category], file)
			case "yellow":
				yellowFiles[category] = append(yellowFiles[category], file)
			case "red":
				redFiles[category] = append(redFiles[category], file)
			}
		}
	}

	// If there are red files, only print those
	if len(redFiles) > 0 {
		fmt.Println("\n❌ Failed")
		printFilesByStatusMap(redFiles, true)
		PrintContributors(result)
		os.Exit(1)
	}

	// Otherwise print all statuses
	if len(greenFiles) > 0 {
		fmt.Println("\n✅ Passed")
		printFilesByStatusMap(greenFiles, false)
	}

	if len(yellowFiles) > 0 {
		fmt.Println("\n⚠️ Warning(s)")
		printFilesByStatusMap(yellowFiles, true)
	}

	// Print Contributors
	PrintContributors(result)

}

func printFilesByStatusMap(filesByCategory map[string][]File, includeReason bool) {
	for category, files := range filesByCategory {
		fmt.Printf("%s files:\n", strings.ToUpper(category))
		for _, file := range files {
			if includeReason && file.Note != "" {
				fmt.Printf("    - %s (Reason: %s)\n", file.Path, file.Note)
			} else {
				fmt.Printf("    - %s\n", file.Path)
			}
		}
	}
}

func PrintContributors(result *SummaryResult) {
	if len(result.Summary.Contributions.Contributors) > 0 {
		fmt.Println("\nContributors")

		if len(result.Summary.Contributions.Contributors) > 25 {
			showContributors, err := input.UserChoice("There are more than 25 contributors, do you want to show all?")
			if err != nil {
				log.Fatal(err)
			}
			if showContributors {
				for _, contributor := range result.Summary.Contributions.Contributors {
					fmt.Printf(" Contributor: %s, with the Email: %s, have done %d commits\n",
						contributor.Name, contributor.Email, contributor.Commits)
				}
			}

		} else {
			for _, contributor := range result.Summary.Contributions.Contributors {
				fmt.Printf(" Contributor: %s, with the Email: %s, have done %d commits\n",
					contributor.Name, contributor.Email, contributor.Commits)
			}
		}

	}
}

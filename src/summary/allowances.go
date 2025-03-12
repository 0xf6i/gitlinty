package summary

import "linty/src/config"

func GetFailureAllowance(config *config.Config, summary *Summary) {
	summary.Readme = config.FailureAllowances.Readme
	summary.License = config.FailureAllowances.License
	summary.Gitignore = config.FailureAllowances.Gitignore
	summary.Workflow = config.FailureAllowances.Workflow
	summary.Tests = config.FailureAllowances.Tests
}

func IsFailureAllowed(category string, summary Summary) bool {
	switch category {
	case "readme":
		return summary.Readme
	case "license":
		return summary.License
	case "gitignore":
		return summary.Gitignore
	case "workflow":
		return summary.Workflow
	case "tests":
		return summary.Tests
	default:
		return false
	}
}

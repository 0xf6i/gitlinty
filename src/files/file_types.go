package files

import (
	"strings"
)

// IsTestFile checks if a file is a test file based on its path and name
func IsTestFile(lowerFileName, lowerPath string) bool {
	// Check for test files with explicit patterns
	if strings.Contains(lowerFileName, "_test.go") ||
		strings.Contains(lowerFileName, "/test") ||
		(strings.Contains(lowerPath, "/test/") && !strings.Contains(lowerPath, "/node_modules/")) {
		return true
	}

	// Check for special test frameworks
	if strings.Contains(lowerPath, "/tests/") ||
		strings.Contains(lowerPath, "/unittest/") ||
		strings.Contains(lowerPath, "/__tests__/") {
		return true
	}

	return false
}

// IsWorkflowFile checks if a file is a workflow file
func IsWorkflowFile(lowerFileName, lowerPath string) bool {
	return strings.Contains(lowerPath, ".github/workflows/") &&
		(strings.HasSuffix(lowerFileName, ".yml") || strings.HasSuffix(lowerFileName, ".yaml"))
}

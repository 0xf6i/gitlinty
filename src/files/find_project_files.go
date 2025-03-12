package files

import (
	"fmt"
	"linty/src/summary"
)

func FindProjectFiles(clonedRepoPath string, ignoredPaths []string, ignoredPatterns []string) ([][]summary.File, error) {
	var allFiles [][]summary.File
	// fmt.Println("cloned repo path:", clonedRepoPath)

	licenseFiles, err := CheckFileContent(clonedRepoPath, "license", ignoredPaths, ignoredPatterns)
	if err != nil {
		return nil, fmt.Errorf("failed to read license files %s", err)
	}
	gitIgnoreFiles, err := CheckFileContent(clonedRepoPath, "gitignore", ignoredPaths, ignoredPatterns)
	if err != nil {
		return nil, fmt.Errorf("failed to read gitignore files %s", err)
	}
	readmeFiles, err := CheckFileContent(clonedRepoPath, "readme", ignoredPaths, ignoredPatterns)
	if err != nil {
		return nil, fmt.Errorf("failed to read readme files %s", err)
	}
	workFlowFiles, err := CheckFileContent(clonedRepoPath, "workflow", ignoredPaths, ignoredPatterns)
	if err != nil {
		return nil, fmt.Errorf("failed to read workflow files %s", err)
	}
	testFiles, err := CheckFileContent(clonedRepoPath, "test", ignoredPaths, ignoredPatterns)
	if err != nil {
		return nil, fmt.Errorf("failed to read test files %s", err)
	}

	allFiles = append(allFiles, licenseFiles)
	allFiles = append(allFiles, gitIgnoreFiles)
	allFiles = append(allFiles, readmeFiles)
	allFiles = append(allFiles, workFlowFiles)
	allFiles = append(allFiles, testFiles)

	return allFiles, nil
}

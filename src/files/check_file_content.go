package files

import (
	"fmt"
	"linty/src/summary"
	"os"
	"path/filepath"
	"strings"
)

func CheckFileContent(dirPath string, typeOfFile string, gitignorePaths []string, gitignorePatterns []string) ([]summary.File, error) {

	var matchedFiles []summary.File
	lowerFileType := strings.ToLower(typeOfFile)

	err := filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error occurred: %s", err)
		}

		if d.IsDir() {
			absPath, err := filepath.Abs(path)
			if err != nil {

				return fmt.Errorf("error occurred: %s", err)
			}
			if IsIgnored(absPath, gitignorePaths, gitignorePatterns) {

				return filepath.SkipDir
			}
			return nil
		}

		absPath, err := filepath.Abs(path)
		if err != nil {
			return fmt.Errorf("error occurred: %s", err)
		}

		lowerFileName := strings.ToLower(d.Name())
		lowerPath := strings.ToLower(path)

		if IsIgnored(absPath, gitignorePaths, gitignorePatterns) {

			return nil
		}
		addFile := func(fileType string) {
			isEmpty, err := IsFileEmpty(absPath)
			if err != nil {

				isEmpty = false
			}
			note := ""
			if isEmpty {
				note = "Empty file"
			}
			matchedFiles = append(matchedFiles, summary.File{Type: fileType, Path: absPath, Note: note})
		}

		switch lowerFileType {
		case "gitignore":
			if lowerFileName == ".gitignore" {
				addFile("gitignore")
			}
		case "readme":
			if lowerFileName == "readme.md" {
				addFile("readme")
			}
		case "license":
			if lowerFileName == "license" {
				addFile("license")
			}
		case "workflow":
			if IsWorkflowFile(lowerFileName, lowerPath) {
				addFile("workflow")
			}
		case "test":
			if IsTestFile(lowerFileName, lowerPath) {
				addFile("test")
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return matchedFiles, nil
}

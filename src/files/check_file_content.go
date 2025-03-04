package files

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CheckFileContent(dirPath string, typeOfFile string) ([]string, error) {
	var matchedFiles []string
	lowerFileType := strings.ToLower(typeOfFile)
	fmt.Println(lowerFileType)

	err := filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return errors.New("errors happen sometimes and its ok")
		}
		if !d.IsDir() {
			lowerFileName := strings.ToLower(d.Name())
			lowerPath := strings.ToLower(path)

			fmt.Println(lowerFileName)

			switch lowerFileType {
			case "gitignore":
				if lowerFileName == ".gitignore" {
					matchedFiles = append(matchedFiles, path)
				}
			case "readme":
				// Ensure "README" is the prefix and check for valid suffixes
				if (strings.HasPrefix(lowerFileName, "readme") || strings.HasPrefix(lowerFileName, "readme.")) &&
					(strings.HasSuffix(lowerFileName, ".md") ||
						strings.HasSuffix(lowerFileName, ".txt") ||
						strings.HasSuffix(lowerFileName, ".rst") ||
						strings.HasSuffix(lowerFileName, ".html")) {
					matchedFiles = append(matchedFiles, path)
				}
			case "license":
				// Check if the file name starts with "license" and matches an acceptable extension
				if (strings.HasPrefix(lowerFileName, "license") ||
					strings.HasPrefix(lowerFileName, "LICENSE")) &&
					(strings.HasSuffix(lowerFileName, ".txt") ||
						strings.HasSuffix(lowerFileName, ".md") ||
						strings.HasSuffix(lowerFileName, ".rst")) {
					matchedFiles = append(matchedFiles, path)
				}
			case "workflow":
				// Ensure the file is within ".github/workflows/" and ends with .yml or .yaml
				if strings.Contains(lowerPath, ".github/workflows/") &&
					(strings.HasSuffix(lowerFileName, ".yml") || strings.HasSuffix(lowerFileName, ".yaml")) {
					matchedFiles = append(matchedFiles, path)
				}
			case "test":
				if strings.Contains(lowerPath, "test") {
					matchedFiles = append(matchedFiles, path)
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matchedFiles, nil
}

// switch lowerFileType {
// case "license":
// 	if

// case "gitignore":
// case "readme":
// }

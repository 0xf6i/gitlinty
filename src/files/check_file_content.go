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

			switch lowerFileType {
			case "gitignore":
				if lowerFileName == ".gitignore" {
					matchedFiles = append(matchedFiles, path)
				}
			case "readme":
				if strings.HasSuffix(lowerFileName, "readme") {
					matchedFiles = append(matchedFiles, path)
				}
			case "workflow":
				if strings.Contains(lowerPath, ".github/workflows/") && (strings.HasSuffix(lowerFileName, ".yml")) || (strings.HasSuffix(lowerFileName, ".yaml")) {
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

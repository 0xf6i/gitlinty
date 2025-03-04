package files

import (
	"fmt"
	"os"
	"strings"
)

func CheckForTest(filePath string) (bool, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	fmt.Println(fileContent)

	fileString := string(filePath)
	if strings.Contains(fileString, "test") {
		return true, nil
	}
	return false, nil

}

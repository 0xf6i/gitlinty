package files

import (
	"bufio"
	"io"
	"os"
	"unicode"
)

func IsFileEmpty(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return false, err
	}
	if fileInfo.Size() == 0 {
		return true, nil
	}

	reader := bufio.NewReader(file)
	hasContent := false

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return false, err
		}

		if !unicode.IsSpace(r) {
			hasContent = true
			break
		}
	}

	return !hasContent, nil
}

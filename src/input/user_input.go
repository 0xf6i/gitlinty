package input

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func UserInput(inputString string) (string, error) {
	fmt.Println(inputString)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if scanner.Text() == "" {
		return "", errors.New("no input was given")
	}
	return scanner.Text(), nil
}

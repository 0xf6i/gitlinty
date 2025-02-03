package input

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func InputReader(inputString string) (string, error) {
	fmt.Println("[INPUT]: STARTED SCANNER")
	fmt.Println(inputString)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	fmt.Println("[INPUT]: GOT INPUT STRING")
	if scanner.Text() == "" {
		return "", errors.New("no input was given")
	}
	fmt.Println("[INPUT]: RETURNING VALUE")
	return scanner.Text(), nil
}

func ChoiceReader(inputString string) (bool, error) {
	fmt.Println(inputString)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	choiceBool := false
	capitalizedInput := strings.ToUpper(scanner.Text())
	if capitalizedInput == "Y" || capitalizedInput == "YES" {
		choiceBool = true
	} else if capitalizedInput == "N" || capitalizedInput == "NO" {
		choiceBool = false
	} else {
		fmt.Println("expected (y/yes) or (n/no) as input")
		return ChoiceReader(inputString)
	}
	fmt.Println(choiceBool)
	return choiceBool, nil
}

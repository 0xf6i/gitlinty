package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UserChoice(inputString string) (bool, error) {
	fmt.Println(inputString + " - (yes/no)")
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
		return UserChoice(inputString)
	}
	return choiceBool, nil
}

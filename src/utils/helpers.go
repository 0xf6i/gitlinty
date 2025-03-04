package utils

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"
)

func GenerateBase64() string {
	// Correct time formatting
	currentTime := time.Now().Format("2006-01-02-15:04:05")
	// Base64 encoding
	stringEncode := base64.StdEncoding.EncodeToString([]byte(currentTime))
	return stringEncode
}

func DecodeBase64(b64string string) string {

	decodedBytes, err := base64.StdEncoding.DecodeString(b64string)
	if err != nil {
		log.Fatal(err)
	}

	decodedString := string(decodedBytes)
	return decodedString
}

func FolderExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println(path, "does not exist")
		return false
	} else {
		fmt.Println("The provided directory named", path, "exists")
		return true
	}
}

package utils

import (
	"encoding/hex"
	"log"
	"time"

	"github.com/google/uuid"
)

func GenerateUuid() string {
	id := uuid.New()
	return id.String()
}

func GenerateBase64() string {
	currentTime := time.Now().Format("2006-01-02 15:05:05.1111")
	stringEncode := hex.EncodeToString([]byte(currentTime))
	return stringEncode
}

func DecodeBase64(hexString string) string {

	decodedBytes, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal(err)
	}

	decodedString := string(decodedBytes)
	return decodedString
}

package utils

import (
	"github.com/google/uuid"
)

func GenerateUuid() string {
	id := uuid.New()
	return id.String()
}

package utils

import (
    "github.com/google/uuid"
)

func GenerateUUIDHex() uuid.UUID {
	return uuid.New()
}
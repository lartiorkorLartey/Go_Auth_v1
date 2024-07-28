package utils

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
)

func GenerateAPN(length int) (string, error) {
    bytes := make([]byte, length)
    if _, err := rand.Read(bytes); err != nil {
        return "", fmt.Errorf("failed to generate random bytes: %w", err)
    }
    return hex.EncodeToString(bytes), nil
}
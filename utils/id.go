package utils

import (
    "crypto/rand"
    "encoding/hex"
)

func GenerateID() string {
    bytes := make([]byte, 4)
    rand.Read(bytes)
    return hex.EncodeToString(bytes)
}
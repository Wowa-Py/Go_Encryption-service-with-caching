package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func HashMD5(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func HashSHA256(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

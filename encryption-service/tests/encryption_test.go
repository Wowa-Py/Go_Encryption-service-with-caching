package tests

import (
	"encryption-service/utils"
	"testing"
)

func TestHashMD5(t *testing.T) {
	input := "test"
	expected := "098f6bcd4621d373cade4e832627b4f6"

	result := utils.HashMD5(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestHashSHA256(t *testing.T) {
	input := "test"
	expected := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"

	result := utils.HashSHA256(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

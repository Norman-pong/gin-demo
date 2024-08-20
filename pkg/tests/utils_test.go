package utils

import (
	"testing"
	"time"

	"zhiming.cool/go/pkg/utils"
)

// TestGenerateRandomString tests the GenerateRandomString function
func TestGenerateRandomString(t *testing.T) {
	length := 10
	str, err := utils.GenerateRandomString(length)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(str) != length*2 {
		t.Fatalf("Expected string length %d, got %d", length*2, len(str))
	}
}

// TestContains tests the Contains function
func TestContains(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}
	if !utils.Contains(slice, "banana") {
		t.Fatalf("Expected true, got false")
	}
	if utils.Contains(slice, "grape") {
		t.Fatalf("Expected false, got true")
	}
}

// TestReverseString tests the ReverseString function
func TestReverseString(t *testing.T) {
	input := "hello"
	expected := "olleh"
	output := utils.ReverseString(input)
	if output != expected {
		t.Fatalf("Expected %s, got %s", expected, output)
	}
}

// TestFormatDate tests the FormatDate function
func TestFormatDate(t *testing.T) {
	layout := "2006-01-02"
	date := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	expected := "2023-10-01"
	output := utils.FormatDate(date, layout)
	if output != expected {
		t.Fatalf("Expected %s, got %s", expected, output)
	}
}

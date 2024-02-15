package strings

import (
	"bytes"
	"strings"
	"testing"
)

// A helper function; This is releated in multiple tests.
func writeTestString(str string, container *StringContainer) {
	for i := 0; i < len(str); i++ {
		container.ReadNextChar(str[i])
	}
}

func TestRead(t *testing.T) {
	test_string := "abcdefg"
	input := strings.NewReader(test_string)
	output := new(bytes.Buffer)
	container := StringContainer{}
	container.Read("\n", input, output)
	result := output.String()
	if result != test_string+"\n" {
		t.Errorf("Expected '%s', got '%s'", test_string+"\n", result)
	}
	// Next test
	container.Reset()
	test_string = "abcdefg" + string(rune(0x19)) + "hijklmnop"
	input = strings.NewReader(test_string)
	output = new(bytes.Buffer)
	container.Read("\n", input, output)
	result = output.String()
	expected_string := "abcdefg\nhijklmnop\n"
	if result != expected_string {
		t.Errorf("Expected '%s', got '%s'", expected_string, result)
	}
}

func TestGetString(t *testing.T) {
	container := StringContainer{}
	test_string := "abcdefg"
	writeTestString(test_string, &container)
	result := container.GetString()
	if result != test_string {
		t.Errorf("Expected %s, got %s", test_string, result)
	}
}

func TestReset(t *testing.T) {
	container := StringContainer{}
	test_string := "abcdefg"
	writeTestString(test_string, &container)
	container.Reset()
	result := container.GetString()
	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}
}

func TestReadNextChar(t *testing.T) {
	container := StringContainer{}
	test_string := "abcdefg"
	writeTestString(test_string, &container)
	// Before reading invalid characters
	result := container.GetString()
	if result != test_string {
		t.Errorf("Expected %s, got %s", test_string, result)
	}
	container.ReadNextChar(0x7F)
	container.ReadNextChar(0x19)
	// After reading invalid characters
	result = container.GetString()
	if result != test_string {
		t.Errorf("Expected %s, got %s", test_string, result)
	}
}

func TestCurrentLength(t *testing.T) {
	container := StringContainer{}
	test_string := "abcdefg"
	for i := 0; i < len(test_string); i++ {
		container.ReadNextChar(test_string[i])
		length := container.GetCurrentLength()
		if length != i+1 {
			t.Errorf("Expected length %d, got lenth %d. Current string: %s", i+1, length, container.GetString())
		}
	}
}

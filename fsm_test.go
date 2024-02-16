package strings

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// A helper function; This is releated in multiple tests.
func writeTestString(str string, container *StringContainer) {
	for i := 0; i < len(str); i++ {
		container.ReadNextChar(str[i])
	}
}

// FIXME: This needs updating from when Read became a higher-order function
func TestRead(t *testing.T) {
	testString := "abcdefg"
	input := strings.NewReader(testString)
	output := new(bytes.Buffer)
	testWriter := func(str string, pos uint64) { fmt.Fprintf(output, "%x %s\n", pos, str) }
	container := StringContainer{}
	container.Read(testWriter, input)
	result := output.String()
	expected := fmt.Sprintf("0 %s\n", testString)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
	// Next test
	container.Reset()
	testString = "abcdefg" + string(rune(0x19)) + "hijklmnop"
	input = strings.NewReader(testString)
	output = new(bytes.Buffer)
	container.Read(testWriter, input)
	result = output.String()
	expected_string := "0 abcdefg\n8 hijklmnop\n"
	if result != expected_string {
		t.Errorf("Expected '%s', got '%s'", expected_string, result)
	}
}

func TestGetString(t *testing.T) {
	container := StringContainer{}
	testString := "abcdefg"
	writeTestString(testString, &container)
	result := container.GetString()
	if result != testString {
		t.Errorf("Expected %s, got %s", testString, result)
	}
}

func TestReset(t *testing.T) {
	container := StringContainer{}
	testString := "abcdefg"
	writeTestString(testString, &container)
	container.Reset()
	result := container.GetString()
	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}
}

func TestReadNextChar(t *testing.T) {
	container := StringContainer{}
	testString := "abcdefg"
	writeTestString(testString, &container)
	// Before reading invalid characters
	result := container.GetString()
	if result != testString {
		t.Errorf("Expected %s, got %s", testString, result)
	}
	container.ReadNextChar(0x7F)
	container.ReadNextChar(0x19)
	// After reading invalid characters
	result = container.GetString()
	if result != testString {
		t.Errorf("Expected %s, got %s", testString, result)
	}
}

func TestCurrentLength(t *testing.T) {
	container := StringContainer{}
	testString := "abcdefg"
	for i := 0; i < len(testString); i++ {
		container.ReadNextChar(testString[i])
		length := container.GetCurrentLength()
		if length != i+1 {
			t.Errorf("Expected length %d, got lenth %d. Current string: %s", i+1, length, container.GetString())
		}
	}
}

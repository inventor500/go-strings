package strings

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	testString := "abcdefg"
	input := strings.NewReader(testString)
	output := new(bytes.Buffer)
	testWriter := func(str string, pos uint64) { fmt.Fprintf(output, "%x %s\n", pos, str) }
	container := StringContainer{}
	container.Read(testWriter, testTestChar, input)
	result := output.String()
	expected := fmt.Sprintf("0 %s\n", testString)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
	container.Reset()
	// That this still workds when an invalid char is added to the mix
	testString = "abcdefg" + string(rune(0x19)) + "hijklmnop"
	input = strings.NewReader(testString)
	output = new(bytes.Buffer)
	container.Read(testWriter, testTestChar, input)
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
	container.ReadNextChar(0x7F, testTestChar)
	container.ReadNextChar(0x19, testTestChar)
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
		container.ReadNextChar(testString[i], testTestChar)
		length := container.GetCurrentLength()
		if length != i+1 {
			t.Errorf("Expected length %d, got lenth %d. Current string: %s", i+1, length, container.GetString())
		}
	}
}

// This is reused in multiple tests
func testTestChar(b byte) bool {
	return b >= 0x20 && b <= 0x7E
}

// A helper function; This is releated in multiple tests.
func writeTestString(str string, container *StringContainer) {
	for i := 0; i < len(str); i++ {
		container.ReadNextChar(str[i], testTestChar)
	}
}

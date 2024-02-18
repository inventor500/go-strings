package strings

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// This funtions like a finite state machine:
// (norma) -> ReadNextChar(ASCII_Char) -> (normal)
// (normal) -> ReadNextChar(not ASCII_Char) -> (ready to reset)
// (ready to reset) -> Reset() -> (normal)
// (ready to reset) -> GetString() -> (ready to reset)
// Use StringContainer.Read to read from an io.Reader and write to an io.Writer
type StringContainer struct {
	Chars  strings.Builder
	Length int
}

func (s *StringContainer) Read(printer func(string, uint64), tester func(byte) bool, r io.Reader) error {
	var position uint64
	reader := bufio.NewReader(r)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			err_str := fmt.Sprintf("%s", err)
			if err_str != "EOF" {
				return err
			}
			break
		}
		// KLUDGE: There should be a better way to track position
		position++
		// If another byte cannot be added to the string, then print (if possible) and reset
		if !s.ReadNextChar(b, tester) { // Byte is invalid
			if s.GetCurrentLength() >= s.Length {
				printer(s.GetString(), position-uint64(s.GetCurrentLength())-1)
			}
			s.Reset()
		}
	}
	if s.GetCurrentLength() >= s.Length {
		// The lack of a '-1' here prevents an off-by-one error
		printer(s.GetString(), position-uint64(s.GetCurrentLength()))
	}
	return nil
}

// Allows for use with fmt (not currently called, though)
func (s StringContainer) String() string {
	return fmt.Sprintf("'%s'; max length: %d", s.GetString(), s.Length)
}

// Get the stored value
func (s StringContainer) GetString() string {
	return s.Chars.String()
}

// Reset the accumulated characters
func (s *StringContainer) Reset() {
	s.Chars.Reset()
}

// Get the current string length stored in the StringContainer
func (s StringContainer) GetCurrentLength() int {
	return s.Chars.Len()
}

// If the byte is a valid part of a string, then add it to the string
// Otherwide, StringContainer is ready to print and reset
func (s *StringContainer) ReadNextChar(b byte, tester func(byte) bool) bool {
	if tester(b) {
		s.Chars.WriteByte(b)
		return true
	}
	return false
}

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
	Chars strings.Builder
	// Position is at the next place in the array to insert.
	Length int
}

func (s *StringContainer) Read(separator string, r io.Reader, w io.Writer) error {
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
		if !s.ReadNextChar(b) {
			if s.GetCurrentLength() > s.Length {
				fmt.Fprintf(w, "%s%s", s.GetString(), separator)
			}
			s.Reset()
		}
	}
	if s.GetCurrentLength() >= s.Length {
		fmt.Fprintf(w, "%s%s", s.GetString(), separator)
	}
	return nil
}

func (s StringContainer) String() string {
	return fmt.Sprintf("'%s'; max length: %d", s.GetString(), s.Length)
}

// Get the stored value
func (s StringContainer) GetString() string {
	// return fmt.Sprintf("%x: %s", s.StartingPosition, s.Chars.String())
	return s.Chars.String()
}

// Reset the accumulated characters
func (s *StringContainer) Reset() {
	s.Chars.Reset()
}

// If the byte is a valid part of a string, then add it to the string and return true
// Otherwide, return false (StringContainer is ready to print and reset)
func (s *StringContainer) ReadNextChar(b byte) bool {
	// Normal ASCII letters
	// Carriage return and line feed do not count
	if b >= 0x20 && b <= 0x7E {
		// if b >= 0x20 && b <= 0x7E {
		s.Chars.WriteByte(b)
		return true
	} else {
		return false
	}
}

func (s StringContainer) GetCurrentLength() int {
	return s.Chars.Len()
}

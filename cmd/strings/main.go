package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	gostr "github.com/inventor500/go-strings"
)

func getArgs() (string, int, string, byte) {
	// 4 is the POSIX-specified default
	length := flag.Int("length", 4, "The minimum length of a string")
	separator := flag.String("output-separator", "\n", "The separator to divide matches")
	radix := flag.String("t", "", "Print the offset before each line. 'x' is hexedecimal, 'o' is octal, and 'd' is decimal. Default is to omit offset.")
	flag.Parse()
	var radix_evaluated byte
	if len(*radix) != 0 {
		radix_evaluated = strings.ToLower(*radix)[0]
	}
	if len(flag.Args()) == 1 {
		return flag.Args()[0], *length, *separator, radix_evaluated
	} else {
		fmt.Fprintf(os.Stderr, "Usage: %s [args] filename\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	// Will never reach here...
	return "", 0, "", 0
}

// Check if POSIXLY_CORRECT or POSIX_ME_HARDER are set
func isPosix() bool {
	val := os.Getenv("POSIXLY_CORRECT")
	isSet, err := strconv.ParseBool(val)
	if err == nil && isSet {
		return true
	}
	val = os.Getenv("POSIX_ME_HARDER")
	isSet, err = strconv.ParseBool(val)
	return err == nil && isSet
}

// Create a writer
func createWriter(separator string, format byte, output *bufio.Writer) func(string, uint64) {
	// This is the only case where the number of required elements changes
	if format == 0 {
		return func(str string, _ uint64) { fmt.Fprintf(output, "%s%s", str, separator) }
	}
	// POSIX requires we not print 0x or 0o
	printIndicator := !isPosix()
	var formatString string
	switch format {
	case 'd':
		formatString = "%d %s%s"
	case 'x':
		if printIndicator {
			formatString = "0x%x %s%s"
		} else {
			formatString = "%x %s%s"
		}
	case 'o':
		if printIndicator {
			formatString = "0o%o %s%s"
		} else {
			formatString = "%o %s%s"
		}
	default:
		fmt.Fprintf(os.Stderr, "Invalid radix %v", format)
		// TODO: Does this leave the input file unclosed?
		os.Exit(1)
	}
	return func(str string, position uint64) {
		fmt.Fprintf(output, formatString, position, str, separator)
	}
}

func main() {
	filename, length, separator, radix := getArgs()
	var file *os.File
	if filename == "-" {
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(filename)
		if err != nil {
			// Error should String() to "open <filename>: <error description>"
			fmt.Fprintf(os.Stderr, "Failed to %s\n", err)
			os.Exit(1)
		}
		// We do not need to close os.Stdout
		defer file.Close()
	}
	var container = gostr.StringContainer{
		Length: length,
	}
	output := bufio.NewWriter(os.Stdout)
	container.Read(createWriter(separator, radix, output), file)
	output.Flush()
	// Make sure that this ends in a new line
	if separator[len(separator)-1] != '\n' {
		fmt.Println()
	}
}

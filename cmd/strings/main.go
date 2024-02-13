package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	gostr "github.com/inventor500/go-strings"
)

var Length int
var Separator string

func getArgs() string {
	length := flag.Int("length", 10, "The minimum length of a string")
	separator := flag.String("output-separator", "\n", "The separator to divide matches. Default is the new line character.")
	flag.Parse()
	if len(flag.Args()) != 1 {
		panic("Requires exactly one file")
	}
	Length = *length
	Separator = *separator
	return flag.Args()[0]
}

func main() {
	filename := getArgs()
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var container gostr.StringContainer
	container.Length = Length
	reader := bufio.NewReader(file)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			err_str := fmt.Sprintf("%s", err)
			if err_str != "EOF" {
				fmt.Printf("Error: %s", err_str)
			}
			break
		}
		if !container.ReadNextChar(b) {
			if container.GetCurrentLength() > container.Length {
				fmt.Printf("%s%s", container.GetString(), Separator)
			}
			container.Reset()
		}
	}
	if container.GetCurrentLength() >= container.Length {
		fmt.Printf("%s%s", container.GetString(), Separator)
	}
}

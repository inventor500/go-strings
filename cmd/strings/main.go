package main

import (
	"flag"
	"fmt"
	"os"

	gostr "github.com/inventor500/go-strings"
)

func getArgs() (string, int, string) {
	length := flag.Int("length", 10, "The minimum length of a string")
	separator := flag.String("output-separator", "\n", "The separator to divide matches")
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [args] filename\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	return flag.Args()[0], *length, *separator
}

func main() {
	filename, length, separator := getArgs()
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var container = gostr.StringContainer{
		Length: length,
	}
	container.Read(separator, file, os.Stdout)
	if separator[len(separator)-1] != '\n' {
		fmt.Println()
	}
}

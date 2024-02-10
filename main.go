package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func getArgs() (int, string) {
	length := flag.Int("length", 10, "The minimum length of a string")
	flag.Parse()
	if len(flag.Args()) != 1 {
		panic("Requires exactly one file")
	}
	return *length, flag.Args()[0]
}

func main() {
	length, filename := getArgs()
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var container StringContainer
	container.Length = length
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
				fmt.Println(container.GetString())
			}
			container.Reset()
		}
	}
	if container.GetCurrentLength() >= container.Length {
		fmt.Println(container.GetString())
	}
}

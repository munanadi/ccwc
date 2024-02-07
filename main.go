package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	flag.Parse()

	// Read a file from the args.
	// go run main.go -- -c test.txt
	// -c
	option := flag.Arg(0)
	// test.txt
	fileName := flag.Arg(1)

	// Open the file
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("%s not found in the current directory", fileName)
	}

	countOfBytes := len(fileBytes)

	lineCount := 0

	// Parse all the bytes and search for line breaks denoted by \n
	for _, v := range fileBytes {
		if v == '\n' {
			lineCount += 1
		}
	}

	words := bytes.Fields(fileBytes)
	wordCount := len(words)

	// Number of words in file

	switch option {
	case "-c":
		fmt.Printf("\t%d\t%s", countOfBytes, fileName)
	case "-l":
		fmt.Printf("\t%d\t%s", lineCount, fileName)
	case "-w":
		fmt.Printf("\t%d\t%s", wordCount, fileName)
	default:
		fmt.Printf("%s that option is not found", option)
	}

}

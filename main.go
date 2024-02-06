package main

import (
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

	switch option {
	case "-c":
		fmt.Printf("\t%d\t%s", countOfBytes, fileName)
	case "-l":
		fmt.Printf("\t%d\t%s", lineCount, fileName)
	}

}

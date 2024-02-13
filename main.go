package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)

	lineFlag := flag.Bool("l", false, "counts lines")
	wordFlag := flag.Bool("w", false, "counts words")
	byteFlag := flag.Bool("c", false, "counts bytes")
	runeFlag := flag.Bool("m", false, "counts runes")

	flag.Parse()
	fileName := flag.Arg(0)

	finalBytes := make(chan []byte)

	if fileName == "" {
		go readFromStdIn(finalBytes)
	} else {
		go goReadFileBytes(fileName, finalBytes)
	}

	var fileBytes []byte

	for v := range finalBytes {
		fileBytes = append(fileBytes, v...)
	}

	// Number of bytes
	countOfBytes := len(fileBytes)

	// Line Count
	lineCount := 0
	// Parse all the bytes and search for line breaks denoted by \n
	for _, v := range fileBytes {
		if v == '\n' {
			lineCount += 1
		}
	}

	// Number of bytes. utf-8 encoded runes
	runes := bytes.Runes(fileBytes)
	characterCount := len(runes)

	// Number of words in file
	words := bytes.Fields(fileBytes)
	wordCount := len(words)

	if *byteFlag && *lineFlag && *wordFlag {
		fmt.Printf("\t%d\t%d\t%d\t%s", lineCount, wordCount, countOfBytes, fileName)
	} else if *byteFlag {
		fmt.Printf("\t%d\t%s", countOfBytes, fileName)
	} else if *lineFlag {
		fmt.Printf("\t%d\t%s", lineCount, fileName)
	} else if *wordFlag {
		fmt.Printf("\t%d\t%s", wordCount, fileName)
	} else if *runeFlag {
		fmt.Printf("\t%d\t%s", characterCount, fileName)
	} else {
		fmt.Printf("option not know")
	}

	wg.Wait()
}

func goReadFileBytes(fileName string, finalBytes chan []byte) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("%s not found in the current directory", fileName)
	}

	finalBytes <- bytes
	close(finalBytes)
	wg.Done()
}

func readFromStdIn(finalBytes chan []byte) {
	reader := bufio.NewReader(os.Stdin)
	readBytes, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalln("something went wrong reading bytes from stdin")
	}
	finalBytes <- readBytes
	close(finalBytes)
	wg.Done()
}

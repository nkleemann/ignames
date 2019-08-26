package main

import (
	"fmt"
	"sync"
)

func main() {
	config := prepareApp()

	var waitgroup sync.WaitGroup
	validNames := make(chan string, config.NumberOfNames)

	opts := PatternOptions{
		firstName: "nik",
		numSeqLen: 2}

	searchForPattern(firstNameNumSeq, opts, validNames, &waitgroup)

	for name := range validNames {
		fmt.Printf("[!] Valid: %s\tURL: %s\n", name, baseURL+name)
	}
}

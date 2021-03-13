package main

import (
	"fmt"
	"sync"
)

func main() {
	config := prepareApp()

	var waitgroup sync.WaitGroup
	validNames := make(chan string, config.NumberOfNames)
	
	// insert your own stuff here
	opts := PatternOptions{
		firstName: "?????",
		wordList: []string{
			"?????"
		},
	}

	searchForPattern(firstNameRandWord, opts, validNames, &waitgroup)

	for name := range validNames {
		fmt.Printf("[!] Valid: %s\tURL: %s\n", name, baseURL+name)
	}
}

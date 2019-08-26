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
		wordList: []string{
			"music",
			"av",
			"tech",
			"poly",
			"art",
			"arts",
			"love",
			"fun",
			"creative",
			"sspace",
		},
	}

	searchForPattern(firstNameRandWord, opts, validNames, &waitgroup)

	for name := range validNames {
		fmt.Printf("[!] Valid: %s\tURL: %s\n", name, baseURL+name)
	}
}

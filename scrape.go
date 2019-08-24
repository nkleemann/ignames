package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

// The instagram base URL
const baseURL = "https://instagram.com/"

// Check if the name is available and if so send it to the channel 'validNames'
func checkName(name string, validNames chan string, waitgroup *sync.WaitGroup) {

	resp, err := http.Get(baseURL + name)

	if err != nil {
		fmt.Printf("[*] Response error on: %s\n", baseURL+name)
		return
	}

	defer resp.Body.Close()
	defer waitgroup.Done()

	fmt.Printf("[*] Trying: %s\t Status: %d\n", name, resp.StatusCode)

	switch resp.StatusCode {
	// Name is available
	case 404:
		validNames <- name

	// Name is taken
	case 200:
		break

	// We made too many requests
	case 429:
		// TODO wait automatically
		fmt.Println("[*] Made too many requests, try again later...")
		os.Exit(1)

	// Instagram down omg
	default:
		fmt.Printf("[*] Unhandled Status: %d for %s", resp.StatusCode, name)
		return
	}
}

// Searche for random usernames and return valid ones in a list
// TODO make: searchForPattern(pattern Pattern) {switch on Pattern enum}
func searchForPatternG(seq []rune, validNames chan string, wg *sync.WaitGroup) {
	// Get all permutations of the given sequence
	allPerms := getPerms(string(seq))

	// Check every name matching Pattern A
	for _, perm := range allPerms {
		wg.Add(1)
		go checkName(perm, validNames, wg)
	}
	wg.Wait()
	close(validNames)
}

// func searchForPattern(pattern Pattern)

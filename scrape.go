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

func searchForPattern(p Pattern, opts PatternOptions, validNames chan string, wg *sync.WaitGroup) {
	// Generate all usernames for the specific pattern
	usernames := generateUserNames(p, opts)

	// Check availability of every name
	for _, name := range usernames {
		wg.Add(1)
		go checkName(name, validNames, wg)
	}
	wg.Wait()
	close(validNames)
}

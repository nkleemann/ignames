package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

//
//	TYPES & CONSTANTS
//

// A generationMethod is a function that generates a string of a given length
type generationMethod func(int) string

// Using a third party service,
const baseURL = "https://instagram.com/"

// const baseURL = "https://insta-node.herokuapp.com/_validate_username?username="

//
//	USERNAME GENERATION
//

// Generate a instagram username of a given length based on a generation method
func newUserName(genMethod generationMethod, length int) string {
	return genMethod(length)
}

// Generate a random string
func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//
//	NETWORKING
//

// Check if the username is available
func available(status int) bool {
	return status == 404
}

func checkName(name string, validNames chan string, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()

	resp, err := http.Get(baseURL + name)

	if err != nil {
		fmt.Println("Response error")
		close(validNames)
		return
	}

	fmt.Printf("[*] Trying: %s\t Status: %d\n", name, resp.StatusCode)

	/*
		// We made too many requests
		if resp.StatusCode == 429 {
			// Wait a few secs and try again after
			fmt.Println("[*] Made too many requests.. waiting 1000 milliseconds", "on: ", name, resp.StatusCode)
			time.Sleep(1000 * time.Millisecond)
			fmt.Println("[*] Trying again!")
		}
	*/

	if available(resp.StatusCode) {
		validNames <- name
	}

	resp.Body.Close()
}

// SearchRandomNames searches for random usernames and returns valid ones in a list
func SearchRandomNames(nameLength int, nameCount int, validNames chan string, waitgroup *sync.WaitGroup) {

	// 3- and 4-letter names are all taken
	if nameLength <= 4 {
		return
	}

	for i := 0; i < nameCount; i++ {
		username := newUserName(randomString, nameLength)
		waitgroup.Add(1)
		go checkName(username, validNames, waitgroup)
	}

	waitgroup.Wait()
	close(validNames)
}

//
//	UTILITY
//

// PrepareApp sets all needed configurations
func PrepareApp() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	PrepareApp()

	numberOfNames := 50
	var waitgroup sync.WaitGroup
	validNames := make(chan string, numberOfNames)

	SearchRandomNames(5, numberOfNames, validNames, &waitgroup)

	for name := range validNames {
		fmt.Printf("[!] Valid: %s\t\tURL: %s\n", name, baseURL+name)
	}
}

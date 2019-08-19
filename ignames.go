package main

import (
	"fmt"
	"math/rand"
	"net/http"
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
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//
//	NETWORKING
//

// Check if the username is valid
func valid(status int) bool {
	return status == 200
}

func checkName(name string, validNames chan string) {

	fmt.Printf("[*] Trying: %s\n", name)

	resp, err := http.Get(baseURL + name)

	if err != nil {
		fmt.Println("Response error")
		close(validNames)
	}

	if !valid(resp.StatusCode) {
		validNames <- name
	}

	resp.Body.Close()
}

//
//	UTILITY
//

// SearchRandomNames searches for random usernames and returns valid ones in a list
func SearchRandomNames(nameLength int, nameCount int, validNames chan string) {

	// 3- and 4-letter names are all taken
	if nameLength <= 4 {
		return
	}

	for i := 0; i < nameCount; i++ {
		username := newUserName(randomString, nameLength)
		go checkName(username, validNames)
	}
}

func main() {
	fmt.Println("ignames v 0.0.1")
	rand.Seed(time.Now().UnixNano())

	validNames := make(chan string)
	SearchRandomNames(5, 10, validNames)

	for name := range validNames {
		fmt.Println(name)
	}
}

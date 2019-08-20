package main

import (
	"math/rand"
	"time"
)

// Config holds flags and values used for ignames
type Config struct {
	NumberOfNames int
}

// Sets all needed configurations and seed the PRNG
func prepareApp() Config {
	rand.Seed(time.Now().UnixNano())

	config := Config{
		NumberOfNames: 50}

	return config
}

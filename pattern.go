/*
	Good usernames have specific patterns.

	We offer the following ones:
		- Pattern A: [firstname] + [numericSequence]  example: luna221
		- Pattern B: [firstname] + [randomWord]       example: maxtravels, laurastar
		- Pattern C: [PatternB]  + [numericSequence]  example: maxtravels9, laurastar200
		- Pattern D: [firstname] + [.] + [verb]       example: jens.codes, jessie.travels
		- Pattern E: n           * [randomWord]       example: mytravelhobby, coolmathgames
		- Pattern F: [PatternE]  + [numericSequence]  example: mytravelhobby4, coolmathgames24
		- Pattern G: [randomPermutation]              example: aural (laura)
		- Pattern H: [leetspeak(PatternA-G)]          example: kyli3j33n3rsb1gg3stst4lk3r
*/

package main

// Pattern defines the different types of usernames we can generate.
type Pattern int

// An enumartion of different username generation Patterns.
const (
	A Pattern = iota
	B Pattern
	C Pattern
	D Pattern
	E Pattern
	F Pattern
	G Pattern
	H Pattern
)

// PatternOptions combine parameters needed to generate the Patterns A-H.
type PatternOptions struct {
	firstName       string   // First name of the user
	numSeqLen       int      // How many numbers in a random numeric sequence
	wordList        []string // A list of random words for Patterns B, E
	verbList        []string // A list of random verbs for Pattern D
	randWordCount   int      // How many random words to combine for Pattern E
	wordToPermutate string   // The word to permutate for Pattern G
	enableL33tSp34k bool     // Enable l33tsp34k for Pattern H
}


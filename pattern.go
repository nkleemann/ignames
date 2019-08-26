/*
	Good usernames have specific patterns.

	We offer the following ones:

		- firstNameNumSeq:
				[firstname] + [numericSequence]
				example: luna221

		- firstNameRandWord
				[firstname] + [randomWord]
				example: maxtravels, laurastar

		- firstNameRandWordNumSeq
				[PatternB]  + [numericSequence]
				example: maxtravels9, laurastar200

		- firstNameDotVerb
				[firstname] + [.] + [verb]
				example: jens.codes, jessie.travels

		- randomWordTimesN
				n           * [randomWord]
				example: mytravelhobby, coolmathgames

		- randomWordTimesNnumSeq
				[PatternE]  + [numericSequence]
				example: mytravelhobby4, coolmathgames24

		- randomPermutation
				[randomPermutation]
				example: aural (laura)

		- leetspeak
				[leetspeak(PatternA-G)]
				 example: kyli3j33n3rsb1gg3stst4lk3r
*/

package main

// Pattern defines the different types of usernames we can generate.
type Pattern int

// An enumartion of different username generation Patterns.
const (
	firstNameNumSeq Pattern = iota
	firstNameRandWord
	firstNameRandWordNumSeq
	firstNameDotVerb
	randomWordTimesN
	randomWordTimesNnumSeq
	randomPermutation
	leetspeak
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

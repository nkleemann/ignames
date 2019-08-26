package main

import "math/rand"

var alphabet = []rune("abcdefghijklmnopqrstuvwxyz")
var numbers = []rune("0123456789")

// Generates all usernames based on a given pattern with options
func generateUserNames(p Pattern, opts PatternOptions) []string {
	var names []string

	switch p {
	// [firstname] + [numericSequence]
	case firstNameNumSeq:
		names = genAllPatternA(opts.firstName, opts.numSeqLen)

	// [firstname] + [randomWord]
	case firstNameRandWord:

	// [PatternB]  + [numericSequence]
	case firstNameRandWordNumSeq:

	// [firstname] + [.] + [verb]
	case firstNameDotVerb:

	// n * [randomWord]
	case randomWordTimesN:

	// [PatternE]  + [numericSequence]
	case randomWordTimesNnumSeq:

	// [randomPermutation]
	case randomPermutation:

	// [leetspeak(PatternA-G)]
	case leetspeak:

	}
	return names
}

func genAllPatternA(firstName string, numSeqLen int) []string {
	var names []string

	allCombs := AllCombinations(numbers, numSeqLen)

	for _, comb := range allCombs {
		names = append(names, firstName+comb)
	}
	return names
}

//
//		HELPER
//

// generate a random sequence choosing from a set of characters
func randomSequence(length int, alphanums []rune) string {

	s := make([]rune, length)
	for i := range s {
		s[i] = alphanums[rand.Intn(len(alphanums))]
	}
	return string(s)
}

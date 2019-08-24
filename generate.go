package main

import "math/rand"

var alphabet = []rune("abcdefghijklmnopqrstuvwxyz")
var numbers = []rune("0123456789")

// Generate a username based on a given pattern with options
func generateUserName(p Pattern, opts PatternOptions) string {
	name := ""

	switch p {
	// [firstname] + [numericSequence]
	case A:
		name = genPatternA(opts.firstName, opts.numSeqLen)
	// [firstname] + [randomWord]
	case B:
	// [PatternB]  + [numericSequence]
	case C:
	// [firstname] + [.] + [verb]
	case D:
	// n * [randomWord]
	case E:
	// [PatternE]  + [numericSequence]
	case F:
	// [randomPermutation]
	case G:
	// [leetspeak(PatternA-G)]
	case H:

	}
	return name
}

/*
func genPatternA(firstName string, numSeqLen int) string {
	return firstName + AllCombinations(numbers, numSeqLen)
}
*/

//
//		HELPER
//

// Get all permutations of a string
func getPerms(str string) []string {
	// base case, for one char, all perms are [char]
	if len(str) == 1 {
		return []string{str}
	}

	// current char
	current := str[0:1]
	// remaining string
	remStr := str[1:]

	// get perms for remaining string
	perms := getPerms(remStr)

	// array to hold all perms of the string based on perms of substring
	allPerms := make([]string, 0)

	// for every perm in the perms of substring
	for _, perm := range perms {
		// add current char at every possible position
		for i := 0; i <= len(perm); i++ {
			newPerm := insertAt(i, current, perm)
			allPerms = append(allPerms, newPerm)
		}
	}
	return allPerms
}

// insert a char in a word
func insertAt(i int, char string, perm string) string {
	start := perm[0:i]
	end := perm[i:len(perm)]
	return start + char + end
}

// First generate an alphanumeric sequence of a given length
// then return every permutation of that sequence.
func allPermsOfRandomSequence(length int, runes []rune) []string {
	randSeq := randomSequence(length, runes)
	return getPerms(randSeq)
}

// generate a random sequence choosing from a set of characters
func randomSequence(length int, alphanums []rune) string {

	s := make([]rune, length)
	for i := range s {
		s[i] = alphanums[rand.Intn(len(alphanums))]
	}
	return string(s)
}

// pick a random word from a wordlist
func randomWordFromList(list []string) string {
	return list[rand.Intn(len(list))]
}

package main

// AllCombinations returns every combination of length n from a source set.
// It is a wrapper for the recursive function 'allCombinationsRec'.

// This function is memory-problematic with sets of len > 20 but we'll
// use it only for numeric sets of len <= 10.
func AllCombinations(source []rune, n int) []string {
	var combs []string
	allCombinationsRec(source, n, "", &combs)
	return combs
}

// The recursive algorithm to find combinations of length n
func allCombinationsRec(source []rune, n int, start string, combs *[]string) {
	if len(start) >= n {
		*combs = append(*combs, start)
	} else {
		for _, c := range source {
			allCombinationsRec(source, n, start+string(c), combs)
		}
	}
}

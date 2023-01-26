package _1arraysAndStrings

import (
	"sort"
)

/*
Given two strings, write a func to decide if on is a permutation of the other.

1. Is this case-sensitive?
2. Does whitespace matter?
*/

// isPermutationSort sorts and compares the given strings.
// It assumes that whitespace and case matters.
func isPermutationSort(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aRunes := []rune(a)
	sort.Slice(aRunes, func(i, j int) bool {
		return aRunes[i] < aRunes[j]
	})

	bRunes := []rune(b)
	sort.Slice(bRunes, func(i, j int) bool {
		return bRunes[i] < bRunes[j]
	})

	return string(aRunes) == string(bRunes)
}

// isPermutationWeight uses a map to count the amount of chars.
func isPermutationWeight(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	lookup := make(map[rune]int)

	for _, char := range a {
		lookup[char]++
	}

	for _, char := range b {
		lookup[char]--

		if lookup[char] < 0 {
			return false
		}
	}

	// for _, v := range lookup {
	// 	if v != 0 {
	// 		return false
	// 	}
	// }

	return true
}

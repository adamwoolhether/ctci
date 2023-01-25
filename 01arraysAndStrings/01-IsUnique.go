package _1arraysAndStrings

import (
	"sort"
)

/*
Implement an algorithm to determine if a string has all unique characters.

What if you cannot use additional data structures?

1. Determine the charset of our string. ASCII or UNICODE
*/

// isUnique1 uses a hashmap.
// O(N), Space: O(N)
func isUnique1(input string) bool {
	lookup := make(map[rune]bool, len(input))

	for _, c := range input {
		if lookup[c] {
			return false
		}

		lookup[c] = true
	}

	return true
}

// isUnique1 uses go's internal sort algorithm
// O(N log N), Space: O(N)
func isUnique2(input string) bool {
	runes := []rune(input)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == runes[i+1] {
			return false
		}
	}

	return true
}

package _1arraysAndStrings

import (
	"sort"
)

/*
Implement an algorithm to determine if a string has all unique characters.

What if you cannot use additional data structures?

1. Determine the charset of our string. ASCII or UNICODE
*/

// isUniqueMap uses a hashmap.
// O(N), Space: O(N)
func isUniqueMap(input string) bool {
	lookup := make(map[rune]bool, len(input))

	for _, c := range input {
		if lookup[c] {
			return false
		}

		lookup[c] = true
	}

	return true
}

// isUniqueSort uses go's internal sort algorithm
// O(N log N), Space: O(log N)
func isUniqueSort(input string) bool {
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

// isUniqueLoops uses nested loops.
// O(N^2), Space: O(1)
func isUniqueLoops(input string) bool {
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes); j++ {
			if i != j && runes[i] == runes[j] {
				return false
			}
		}
	}

	return true
}

// isUniqueBitSet
func isUniqueBitSet(input string) bool {
	check := 0
	for _, char := range []rune(input) {
		val := char - 'a'
		mask := 1 << uint32(val)
		if check&mask > 0 {
			return false
		}
		check |= mask
	}

	return true
}

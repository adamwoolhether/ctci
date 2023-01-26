package _1arraysAndStrings

import (
	"testing"
)

/*
gotest -v -run TestIsPermutation
go test -run '^$' -bench '^BenchmarkIsPermutation*'

goos: darwin
goarch: amd64
pkg: github.com/adamwoolhether/ctci/01arraysAndStrings
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkIsPermutationSort-16             329455              3205 ns/op             656 B/op          8 allocs/op
BenchmarkIsPermutationWeight-16           374898              3064 ns/op             830 B/op          5 allocs/op
PASS
ok      github.com/adamwoolhether/ctci/01arraysAndStrings       2.627s
*/

type permutationTest struct {
	name     string
	aString  string
	bString  string
	expected bool
}

var testCases = []permutationTest{
	{
		name:     "dog god",
		aString:  "dog",
		bString:  "god",
		expected: true,
	},
	{
		name:     "dog god",
		aString:  "dogdog",
		bString:  "god",
		expected: false,
	},
	{
		name:     "dog God",
		aString:  "dog",
		bString:  "God",
		expected: false,
	},
	{
		name:     "alphabet all true",
		aString:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		bString:  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
		expected: true,
	},
	{
		name:     "alphabet all false",
		aString:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		bString:  "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ",
		expected: false,
	},
}

func TestIsPermutationSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isPermutationSort(tc.aString, tc.bString)
			if got != tc.expected {
				t.Errorf("got: %t, exp: %t", got, tc.expected)
			}
		})
	}
}

func BenchmarkIsPermutationSort(b *testing.B) {
	aStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = isPermutationSort(aStr, bStr)
	}
}

func TestIsPermutationWeight(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isPermutationWeight(tc.aString, tc.bString)
			if got != tc.expected {
				t.Errorf("got: %t, exp: %t", got, tc.expected)
			}
		})
	}
}

func BenchmarkIsPermutationWeight(b *testing.B) {
	aStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = isPermutationWeight(aStr, bStr)
	}
}

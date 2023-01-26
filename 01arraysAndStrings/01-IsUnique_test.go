package _1arraysAndStrings

import "testing"

/*
gotest -v -run TestIsUnique
go test -run '^$' -bench '^BenchmarkIsUnique*'

goos: darwin
goarch: amd64
pkg: github.com/adamwoolhether/ctci/01arraysAndStrings
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkIsUniqueMap-16           657664              1837 ns/op            1050 B/op          2 allocs/op
BenchmarkIsUniqueSort-16         1000000              1059 ns/op             280 B/op          3 allocs/op
BenchmarkIsUniqueLoops-16        8753114               139.2 ns/op           224 B/op          1 allocs/op
BenchmarkIsUniqueBitSet-16       7581468               161.0 ns/op           224 B/op          1 allocs/op
PASS
ok      github.com/adamwoolhether/ctci/01arraysAndStrings       6.310s


*/

func TestIsUniqueMap(t *testing.T) {
	testCases := []struct {
		input string
		exp   bool
	}{
		{"1àha漢字Pépy5", true},
		{"1àha漢字Pépy51àha漢字Pépy5", false},
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", true},
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZa", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := isUniqueMap(tc.input)
			if got != tc.exp {
				t.Errorf("got %t, exp %t", got, tc.exp)
			}
		})
	}
}

func BenchmarkIsUniqueMap(b *testing.B) {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZa"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isUniqueMap(str)
	}
}

func TestIsUniqueSort(t *testing.T) {
	testCases := []struct {
		input string
		exp   bool
	}{
		{"1àha漢字Pépy5", true},
		{"1àha漢字Pépy51àha漢字Pépy5", false},
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", true},
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZa", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := isUniqueSort(tc.input)
			if got != tc.exp {
				t.Errorf("got %t, exp %t", got, tc.exp)
			}
		})
	}
}

func BenchmarkIsUniqueSort(b *testing.B) {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZa"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isUniqueSort(str)
	}
}

func TestIsUniqueLoops(t *testing.T) {
	testCases := []struct {
		input string
		exp   bool
	}{
		{"1àha漢字Pépy5", true},
		{"1àha漢字Pépy51àha漢字Pépy5", false},
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", true},
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZa", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := isUniqueLoops(tc.input)
			if got != tc.exp {
				t.Errorf("got %t, exp %t", got, tc.exp)
			}
		})
	}
}

func BenchmarkIsUniqueLoops(b *testing.B) {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZa"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isUniqueLoops(str)
	}
}

func TestIsUniqueBitSet(t *testing.T) {
	testCases := []struct {
		input string
		exp   bool
	}{
		{"1àha漢字Pépy5", true},
		{"1àha漢字Pépy51àha漢字Pépy5", false},
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", true},
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZa", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := isUniqueBitSet(tc.input)
			if got != tc.exp {
				t.Errorf("got %t, exp %t", got, tc.exp)
			}
		})
	}
}

func BenchmarkIsUniqueBitSet(b *testing.B) {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZa"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isUniqueBitSet(str)
	}
}

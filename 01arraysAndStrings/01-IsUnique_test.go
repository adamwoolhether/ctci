package _1arraysAndStrings

import "testing"

/*
gotest -v -run TestIsUnique
go test -run '^$' -bench '^BenchmarkIsUnique*'

goos: darwin
goarch: amd64
pkg: github.com/adamwoolhether/ctci/01arraysAndStrings
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkIsUnique1-16             395823              2764 ns/op            1050 B/op          2 allocs/op
BenchmarkIsUnique2-16             707449              1664 ns/op             280 B/op          3 allocs/op
BenchmarkIsUnique3-16            5706530               209.1 ns/op           224 B/op          1 allocs/op
PASS
ok      github.com/adamwoolhether/ctci/01arraysAndStrings       3.902s

*/

func TestIsUnique1(t *testing.T) {
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
			got := isUnique1(tc.input)
			if got != tc.exp {
				t.Errorf("got %t, exp %t", got, tc.exp)
			}
		})
	}
}

func BenchmarkIsUnique1(b *testing.B) {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZa"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isUnique1(str)
	}
}

func TestIsUnique2(t *testing.T) {
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
			got := isUnique2(tc.input)
			if got != tc.exp {
				t.Errorf("got %t, exp %t", got, tc.exp)
			}
		})
	}
}

func BenchmarkIsUnique2(b *testing.B) {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZa"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isUnique2(str)
	}
}

func TestIsUnique3(t *testing.T) {
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
			got := isUnique3(tc.input)
			if got != tc.exp {
				t.Errorf("got %t, exp %t", got, tc.exp)
			}
		})
	}
}

func BenchmarkIsUnique3(b *testing.B) {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZa"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isUnique3(str)
	}
}

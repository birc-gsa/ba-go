package gsa_test

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"birc.au.dk/gsa"
)

// NewRandomSeed creates a new random number generator
func NewRandomSeed(tb testing.TB) *rand.Rand {
	tb.Helper()

	seed := time.Now().UTC().UnixNano()
	// maybe enable this again if it is useful,
	// but right now I don't want it in the benchmarks
	// tb.Logf("Random seed: %d", seed)
	return rand.New(rand.NewSource(seed))
}

// RandomStringN constructs a random string of length in n, over the alphabet alpha.
func RandomStringN(n int, alpha string, rng *rand.Rand) string {
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = alpha[rng.Intn(len(alpha))]
	}

	return string(bytes)
}

// RandomStringRange constructs a random string of length in [min, max), over the alphabet alpha.
func RandomStringRange(min, max int, alpha string, rng *rand.Rand) string {
	n := min + rng.Intn(max-min)
	return RandomStringN(n, alpha, rng)
}

// FibonacciString returns the n'th Fibonacci string.
func FibonacciString(n int) string {
	const (
		fib0 = "a"
		fib1 = "b"
	)

	switch n {
	case 0:
		return fib0

	case 1:
		return fib1

	default:
		a, b := fib0, fib1
		for ; n > 1; n-- {
			a, b = b, a+b
		}

		return b
	}
}

// SingletonString generates a string of length n consisting only of the letter a
func SingletonString(n int, a byte) string {
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = a
	}

	return string(bytes)
}

// PickRandomPrefix returns a random prefix of the string x.
func PickRandomPrefix(x string, rng *rand.Rand) string {
	return x[:rng.Intn(len(x))]
}

// PickRandomSuffix returns a random sufix of the string x.
func PickRandomSuffix(x string, rng *rand.Rand) string {
	return x[rng.Intn(len(x)):]
}

// PickRandomSubstring returns a random substring of the string x.
func PickRandomSubstring(x string, rng *rand.Rand) string {
	i := rng.Intn(len(x) - 1)
	j := rng.Intn(len(x) - i)

	return x[i : i+j]
}

// GenerateRandomTestStrings generates strings of length between
// min and max and calls callback with them.
func GenerateRandomTestStrings(
	min, max int,
	rng *rand.Rand,
	callback func(x string)) {
	n := 50 // number of random strings (maybe parameterise)
	for i := 0; i < n; i++ {
		callback(RandomStringRange(min, max, "abcdefg", rng))
	}
}

// GenerateSingletonTestStrings generate singeton strings with length
// between min and max
func GenerateSingletonTestStrings(
	min, max int,
	rng *rand.Rand,
	callback func(x string)) {
	n := 50 // number of random strings (maybe parameterise)
	for i := 0; i < n; i++ {
		// maybe it is a little overkill to generate this many
		// singletons?
		callback(SingletonString(min+rng.Intn(max-min), 'a'))
	}
}

// GenerateTestStrings generates strings of length between min
// and max and calls callback with them.
func GenerateTestStrings(
	min, max int,
	rng *rand.Rand,
	callback func(x string)) {
	GenerateRandomTestStrings(min, max, rng, callback)
	GenerateSingletonTestStrings(min, max, rng, callback)

	for n := 0; n < 10; n++ {
		callback(FibonacciString(n))
	}
}

func Test_BorderarrayBasics(t *testing.T) {
	tests := []struct {
		name string
		x    string
		want []int
	}{
		{"(empty string)", "", []int{}},
		{"a", "a", []int{0}},
		{"aaa", "aaa", []int{0, 1, 2}},
		{"aaaba", "aaaba", []int{0, 1, 2, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gsa.Borderarray(tt.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Borderarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StrictBorderarrayBasics(t *testing.T) {
	tests := []struct {
		name string
		x    string
		want []int
	}{
		{"(empty string)", "", []int{}},
		{"a", "a", []int{0}},
		{"aaa", "aaa", []int{0, 0, 2}},
		{"aaaba", "aaaba", []int{0, 0, 2, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gsa.StrictBorderarray(tt.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrictBorderarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func checkBorders(t *testing.T, x string, ba []int) {
	t.Helper()

	for i, b := range ba {
		if b != 0 && x[:b] != x[i-b+1:i+1] {
			t.Errorf(`x[:%d] == %q is not a border of %q`, b, x[:b], x[:i+1])
			t.Fatalf(`x = %q, ba = %v`, x, ba)
		}
	}
}

func Test_Borderarray(t *testing.T) {
	rng := NewRandomSeed(t)
	GenerateTestStrings(10, 20, rng,
		func(x string) {
			checkBorders(t, x, gsa.Borderarray(x))
		})
}

func checkStrict(t *testing.T, x string, ba []int) bool {
	t.Helper()

	for i, b := range ba[:len(ba)-1] {
		if b > 0 && x[b] == x[i+1] {
			t.Errorf(`x[:%d] == %q[%q] is not a strict border of %q[%q]`, b, x[:b], x[b], x[:i+1], x[i+1])
			t.Errorf(`x[%d] == %q == x[%d+1] (should be different)`, b, x[b], i)
			t.Fatalf(`x = %q, ba = %v`, x, ba)

			return false
		}
	}

	return true
}

func Test_StrictBorderarray(t *testing.T) {
	rng := NewRandomSeed(t)
	GenerateTestStrings(10, 20, rng,
		func(x string) {
			ba := gsa.StrictBorderarray(x)
			checkBorders(t, x, ba)
			checkStrict(t, x, ba)
		})
}
